package model

import (
	"github.com/filedrive-team/go-filedag-sdk/common"
)

type SearchPinParam struct {
	Cid    []string             `json:"cid,omitempty"`
	Name   string               `json:"name,omitempty"`
	Match  TextMatchingStrategy `json:"match"`
	Status []Status             `json:"status"`
	Before *common.UTCTime      `json:"before,omitempty"`
	After  *common.UTCTime      `json:"after,omitempty"`
	Meta   map[string]string    `json:"meta,omitempty"`
	Limit  int                  `json:"limit"`
	Page   int                  `json:"page,omitempty"` // Paging if set it
}
