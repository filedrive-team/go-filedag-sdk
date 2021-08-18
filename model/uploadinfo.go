package model

type UploadParam struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Mime string `json:"mime"`
}

type UploadResponse struct {
	Token string `json:"token"`
	Host  string `json:"host"`
}

type PinFileResponse struct {
	Cid       string `json:"cid"`
	RequestId string `json:"request_id"`
}
