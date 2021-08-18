package model

type Scope struct {
	Admin     *bool      `json:"admin,omitempty"`
	Endpoints *Endpoints `json:"endpoints,omitempty"`
}

type Endpoints struct {
	Pins    ScopePins    `json:"pins"`
	Data    ScopeData    `json:"data"`
	Pinning ScopePinning `json:"pinning"`
}

type ScopePins struct {
	AddPinObject     bool `json:"addPinObject"`
	GetPinObject     bool `json:"getPinObject"`
	ListPinObjects   bool `json:"listPinObjects"`
	RemovePinObject  bool `json:"removePinObject"`
	ReplacePinObject bool `json:"replacePinObject"`
}

type ScopeData struct {
	PinnedTotal bool `json:"pinnedTotal"`
	UploadFile  bool `json:"uploadFile"`
}

type ScopePinning struct {
	MetaData      bool `json:"metaData"`
	PinPolicy     bool `json:"pinPolicy"`
	UserPinPolicy bool `json:"userPinPolicy"`
}
