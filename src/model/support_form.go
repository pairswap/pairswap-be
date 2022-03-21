package model

type SupportForm struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	TxURL   string `json:"tx_url,omitempty"`
	Comment string `json:"comment,omitempty"`
}
