package model

type PaginationParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
