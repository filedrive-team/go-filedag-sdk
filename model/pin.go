package model

// Pin - Pin object
type Pin struct {

	// Content Identifier (CID) to be pinned recursively
	Cid string `json:"cid" example:"QmCIDToBePinned"`

	// Optional name for pinned data; can be used for lookups later
	Name string `json:"name,omitempty" maximum:"255" example:"PreciousData.pdf"`

	// Optional list of multiaddrs known to provide the data
	Origins []string `json:"origins,omitempty" example:"/ip4/203.0.113.142/tcp/4001/p2p/QmSourcePeerId,/ip4/203.0.113.114/udp/4001/quic/p2p/QmSourcePeerId"`

	// Optional metadata for pin object
	Meta map[string]string `json:"meta,omitempty" example:"app_id:99986338-1113-4706-8302-4420da6158aa"`
}
