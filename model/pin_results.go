package model

type PinResults struct {

	// The total number of pin objects that exist for passed query filters
	Count int32 `json:"count"`

	// An array of object results
	Results []*PinStatus `json:"results"`
}
