package models

// ContaAPagarIncluido estrutura do webhook
type OmieWebhook struct {
	Messageid string `json:"messageId"`
	Topic     string `json:"topic"`
	Author    struct {
		Email  string `json:"email"`
		Name   string `json:"name"`
		Userid int    `json:"userId"`
	} `json:"author"`
	Appkey  string `json:"appKey"`
	Apphash string `json:"appHash"`
	Origin  string `json:"origin"`
}
