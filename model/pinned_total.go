package model

import "github.com/shopspring/decimal"

type PinnedTotal struct {
	Count                     int64           `json:"count"`
	SizeTotal                 decimal.Decimal `json:"size_total"`
	SizeWithReplicationsTotal decimal.Decimal `json:"size_with_replications_total"`
}
