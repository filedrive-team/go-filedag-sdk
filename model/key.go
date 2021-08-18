package model

type ApiKey struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Jwt       string `json:"jwt"`
}
