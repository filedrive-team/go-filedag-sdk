package model

type KeyInfo struct {
	KeyName string `json:"key_name" example:"mykey"`
	MaxUses uint64 `json:"max_uses,omitempty"`
	Scope   Scope  `json:"scope"`
}
