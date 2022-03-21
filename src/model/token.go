package model

type Token struct {
	Id       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	LogoURL  string `json:"logo_url,omitempty"`
	Decimals int    `json:"decimals,omitempty"`
}
