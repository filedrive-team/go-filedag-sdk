package model

import "github.com/filedrive-team/go-filedag-sdk/common"

// PinStatus - Pin object with status
type PinStatus struct {

	// Globally unique identifier of the pin request; can be used to check the status of ongoing pinning, or pin removal
	Requestid string `json:"requestid" example:"UniqueIdOfPinRequest"`

	Status Status `json:"status" example:"queued"`

	// Immutable timestamp indicating when a pin request entered a pinning service; can be used for filtering results and pagination
	Created common.UTCTime `json:"created" example:"2020-07-27T17:32:28Z"`

	Pin Pin `json:"pin"`

	// List of multiaddrs designated by pinning service for transferring any new data from external peers
	Delegates []string `json:"delegates" example:"/ip4/203.0.113.1/tcp/4001/p2p/QmServicePeerId"`

	// Optional info for PinStatus response
	Info map[string]string `json:"info,omitempty" example:"status_details:Queue position 7 of 9"`
}
