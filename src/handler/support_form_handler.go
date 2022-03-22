package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sisu-network/pairswap-be/src/model"
	"github.com/sisu-network/pairswap-be/src/store"
)

type SupportFormRequest struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	TxURL   string `json:"tx_url,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type SupportFormHandler struct {
	db *store.DBStores
}

func NewSupportFormHandler(db *store.DBStores) *SupportFormHandler {
	return &SupportFormHandler{db: db}
}

func (h *SupportFormHandler) HandleSubmitSupportForm(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		_ = responseError(w, err)
		return
	}

	request := &SupportFormRequest{}
	if err := json.Unmarshal(body, request); err != nil {
		fmt.Println(err)
		_ = responseError(w, err)
		return
	}

	supportForm := &model.SupportForm{
		Name:    request.Name,
		Email:   request.Email,
		TxURL:   request.TxURL,
		Comment: request.Comment,
	}

	if err := h.db.SupportFormStore.CreateSupportForm(context.Background(), supportForm); err != nil {
		_ = responseError(w, err)
		return
	}

	if err := request.Send(); err != nil {
		_ = responseError(w, err)
		return
	}

	_ = responseSuccess(w, map[string]interface{}{"msg": "create support form successfully"})
}
