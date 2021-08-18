package client

import "github.com/filedrive-team/go-filedag-sdk/model"

type ResultBase struct {
	Code    int
	Failure *model.Failure
}

func (rb *ResultBase) SetCode(code int) {
	rb.Code = code
}

func (rb *ResultBase) CreateFailure() *model.Failure {
	rb.Failure = &model.Failure{}
	return rb.Failure
}

func (rb *ResultBase) CreateData() interface{} {
	return nil
}

type Responder interface {
	SetCode(code int)
	CreateFailure() *model.Failure
	CreateData() interface{}
}

type ListPinResp struct {
	ResultBase
	Data *model.PinResults
}

func (r *ListPinResp) CreateData() interface{} {
	r.Data = &model.PinResults{}
	return r.Data
}

type PinResp struct {
	ResultBase
	Data *model.PinStatus
}

func (r *PinResp) CreateData() interface{} {
	r.Data = &model.PinStatus{}
	return r.Data
}

type KeyResp struct {
	ResultBase
	Data *model.ApiKey
}

func (r *KeyResp) CreateData() interface{} {
	r.Data = &model.ApiKey{}
	return r.Data
}

type PinnedTotalResp struct {
	ResultBase
	Data *model.PinnedTotal
}

func (r *PinnedTotalResp) CreateData() interface{} {
	r.Data = &model.PinnedTotal{}
	return r.Data
}

type UploadTokenResp struct {
	ResultBase
	Data *model.UploadResponse
}

func (r *UploadTokenResp) CreateData() interface{} {
	r.Data = &model.UploadResponse{}
	return r.Data
}

type PinFileResp struct {
	ResultBase
	Data *model.PinFileResponse
}

func (r *PinFileResp) CreateData() interface{} {
	r.Data = &model.PinFileResponse{}
	return r.Data
}
