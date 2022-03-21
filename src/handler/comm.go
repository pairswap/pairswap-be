package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sisu-network/lib/log"
)

func responseError(w http.ResponseWriter, err error) error {
	resp := DefaultResponse{
		ErrorMsg: err.Error(),
		Data:     nil,
	}

	r, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
		return err
	}

	w.WriteHeader(http.StatusBadRequest)
	if _, err = w.Write(r); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func responseSuccess(w http.ResponseWriter, data map[string]interface{}) error {
	resp := DefaultResponse{
		Data: data,
	}

	r, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
		return err
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(r); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

type DefaultResponse struct {
	ErrorMsg string                 `json:"msg,omitempty"`
	Data     map[string]interface{} `json:"data"`
}
