package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sisu-network/lib/log"
)

const UpdateGasCostInterval = 5 * time.Minute

type GasCostRequest struct {
	TokenId string `json:"token_id,omitempty"`
	Chain   string `json:"chain,omitempty"`
}

// GasCostResponse we use this struct to parse the response from Sisu's server and return it to caller
type GasCostResponse struct {
	Chain   string `json:"chain,omitempty"`
	TokenId string `json:"token_id,omitempty"`
	GasCost int64  `json:"gas_cost,omitempty"`
}

type GasCostRecord struct {
	Cost              int64
	LatestUpdatedTime time.Time
}

type GasCostHandler struct {
	// key: chain + token id
	// value: gas cost record
	GasCostMap map[string]*GasCostRecord

	SisuServerURL   string
	SisuGasCostPath string
}

func NewGasCostHandler(sisuServerURL, sisuGasCostPath string) *GasCostHandler {
	return &GasCostHandler{
		GasCostMap:      make(map[string]*GasCostRecord),
		SisuServerURL:   sisuServerURL,
		SisuGasCostPath: sisuGasCostPath,
	}
}

func (h *GasCostHandler) HandleGetGasCost(w http.ResponseWriter, r *http.Request) {
	queryStr := r.URL.Query()
	tokenId := queryStr.Get("token_id")
	if len(tokenId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("missing token_id")); err != nil {
			log.Warn(err)
		}
		return
	}
	chainId := queryStr.Get("chain")
	if len(chainId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("missing chain")); err != nil {
			log.Warn(err)
		}
		return
	}

	key := h.getKey(chainId, tokenId)
	if record, ok := h.GasCostMap[key]; ok {
		// if not expired then just return the cache record
		if record.LatestUpdatedTime.Add(UpdateGasCostInterval).After(time.Now()) {
			_ = responseSuccess(w, map[string]interface{}{
				"chain":    chainId,
				"token_id": tokenId,
				"gas_cost": record.Cost,
			})

			return
		}
	}

	cost, err := h.getGasCostFromSisu(chainId, tokenId)
	if err != nil {
		log.Error(err)
		_ = responseError(w, err)
		return
	}

	h.GasCostMap[key] = &GasCostRecord{
		Cost:              cost,
		LatestUpdatedTime: time.Now(),
	}

	_ = responseSuccess(w, map[string]interface{}{
		"chain":    chainId,
		"token_id": tokenId,
		"gas_cost": cost,
	})
}

func (h *GasCostHandler) getGasCostFromSisu(chain, token string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, h.SisuServerURL+h.SisuGasCostPath, nil)
	if err != nil {
		log.Error("error when new request ", err)
		return -1, err
	}

	q := req.URL.Query()
	q.Add("token_id", token)
	q.Add("chain", chain)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return -1, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return -1, err
	}

	defer resp.Body.Close()

	gasCostResp := &GasCostResponse{}
	if err := json.Unmarshal(body, gasCostResp); err != nil {
		log.Error(err)
		return -1, err
	}

	return gasCostResp.GasCost, nil
}

func (h *GasCostHandler) getKey(chain, token string) string {
	return fmt.Sprintf("%s__%s", chain, token)
}
