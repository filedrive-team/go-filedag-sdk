package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/filedrive-team/go-filedag-sdk/common"
	"github.com/filedrive-team/go-filedag-sdk/model"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	uploadEndpoint = "/v0/filedag/dn/upload"

	ZH XLanguage = "zh"
	EN XLanguage = "en"
)

type XLanguage string

type Client struct {
	cli       *http.Client
	jwt       string
	apiKey    string
	apiSecret string
	host      string
	language  XLanguage
}

func New() *Client {
	return &Client{
		cli: &http.Client{},
	}
}

func NewWithJwtToken(host string, jwt string) *Client {
	return &Client{
		cli:  &http.Client{},
		jwt:  jwt,
		host: host,
	}
}

func NewWithKeySecret(host string, apiKey, apiSecret string) *Client {
	return &Client{
		cli:       &http.Client{},
		apiKey:    apiKey,
		apiSecret: apiSecret,
		host:      host,
	}
}

func (c *Client) SetHost(host string) {
	c.host = host
}

func (c *Client) SetJwtToken(jwt string) {
	c.jwt = jwt
}

func (c *Client) SetKeySecret(apiKey, apiSecret string) {
	c.apiKey = apiKey
	c.apiSecret = apiSecret
}

func (c *Client) SetLanguage(lang XLanguage) {
	c.language = lang
}

func (c *Client) newRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.host+endpoint, body)
	if err != nil {
		return nil, err
	}
	if len(c.jwt) > 0 {
		req.Header.Add("Authorization", "Bearer "+c.jwt)
	} else {
		req.Header.Add("X-Api-Key", c.apiKey)
		req.Header.Add("X-Api-Secret", c.apiSecret)
	}
	switch c.language {
	case ZH, EN:
		req.Header.Add("X-Language", string(c.language))
	}
	return req, nil
}

func (c *Client) processRequest(method string, endpoint string, reqBodyParam interface{}, responder Responder) (err error) {
	var reqBody io.Reader
	if reqBodyParam != nil {
		var data []byte
		data, err = json.Marshal(reqBodyParam)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(data)
		reqBody = buf
	}
	req, err := c.newRequest(method, endpoint, reqBody)
	if err != nil {
		return err
	}
	resp, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	body := resp.Body
	defer body.Close()

	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	responder.SetCode(resp.StatusCode)
	switch resp.StatusCode {
	case http.StatusOK, http.StatusAccepted:
		if data := responder.CreateData(); data != nil {
			if err = json.Unmarshal(payload, data); err != nil {
				return err
			}
		}
	default:
		failure := responder.CreateFailure()
		if err = json.Unmarshal(payload, &failure); err != nil {
			return err
		}
	}
	return
}

func (c *Client) ListPin(param *model.SearchPinParam) (lpr *ListPinResp, err error) {
	endpoint := "/pins"
	if param != nil {
		query, err := common.ToQueryByJsonTag(param)
		if err != nil {
			return nil, err
		}
		endpoint += "?" + query
	}
	lpr = new(ListPinResp)
	err = c.processRequest(http.MethodGet, endpoint, nil, lpr)
	return
}

func (c *Client) AddPin(pin *model.Pin) (pr *PinResp, err error) {
	endpoint := "/pins"
	pr = new(PinResp)
	err = c.processRequest(http.MethodPost, endpoint, pin, pr)
	return
}

func (c *Client) GetPin(requestId string) (pr *PinResp, err error) {
	endpoint := "/pins/" + requestId
	pr = new(PinResp)
	err = c.processRequest(http.MethodGet, endpoint, nil, pr)
	return
}

func (c *Client) ReplacePin(requestId string, pin *model.Pin) (pr *PinResp, err error) {
	endpoint := "/pins/" + requestId
	pr = new(PinResp)
	err = c.processRequest(http.MethodPost, endpoint, pin, pr)
	return
}

func (c *Client) RemovePin(requestId string) (res *ResultBase, err error) {
	endpoint := "/pins/" + requestId
	res = new(ResultBase)
	err = c.processRequest(http.MethodDelete, endpoint, nil, res)
	return
}

func (c *Client) GenerateApiKey(ki *model.KeyInfo) (kr *KeyResp, err error) {
	endpoint := "/keys"
	kr = new(KeyResp)
	err = c.processRequest(http.MethodPost, endpoint, ki, kr)
	return
}

func (c *Client) RevokeApiKey(apiKey string) (res *ResultBase, err error) {
	endpoint := "/keys/" + apiKey
	res = new(ResultBase)
	err = c.processRequest(http.MethodDelete, endpoint, nil, res)
	return
}

func (c *Client) PinnedDataTotal() (ptr *PinnedTotalResp, err error) {
	endpoint := "/data/pinnedTotal"
	ptr = new(PinnedTotalResp)
	err = c.processRequest(http.MethodGet, endpoint, nil, ptr)
	return
}

func (c *Client) PinFileFromStream(file io.Reader, param *model.UploadParam) (pfr *PinFileResp, err error) {
	if param == nil {
		param = &model.UploadParam{}
	}
	if len(param.Mime) == 0 {
		param.Mime = "application/octet-stream"
	}
	var bodyBuffer *bytes.Buffer
	if param.Size > 0 {
		bodyBuffer = bytes.NewBuffer(make([]byte, param.Size))
		bodyBuffer.Reset()
	} else {
		bodyBuffer = &bytes.Buffer{}
	}
	writer := multipart.NewWriter(bodyBuffer)
	part, err := writer.CreateFormFile("file", "f_"+param.Name)
	if err != nil {
		writer.Close()
		return nil, err
	}

	filesize, err := io.Copy(part, file)
	if err != nil {
		writer.Close()
		return nil, err
	}
	contentType := writer.FormDataContentType()
	writer.Close()

	if param.Size == 0 {
		param.Size = filesize
	}

	utr, err := c.getUploadToken(param)
	if err != nil {
		return nil, err
	}
	if utr.Code != http.StatusOK {
		return nil, errors.New(fmt.Sprint(utr.Failure.Error.Details))
	}

	req, err := http.NewRequest(http.MethodPost, utr.Data.Host+uploadEndpoint, bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+utr.Data.Token)
	req.Header.Add("Content-Type", contentType)

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()

	payload, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	pfr = new(PinFileResp)
	pfr.Code = resp.StatusCode
	switch pfr.Code {
	case http.StatusOK:
		pfr.Data = new(model.PinFileResponse)
		if err = json.Unmarshal(payload, pfr.Data); err != nil {
			return nil, err
		}
	default:
		if err = json.Unmarshal(payload, pfr.Failure); err != nil {
			return nil, err
		}
	}
	return
}

func (c *Client) getUploadToken(param *model.UploadParam) (utr *UploadTokenResp, err error) {
	endpoint := "/data/uploadToken"
	utr = new(UploadTokenResp)
	err = c.processRequest(http.MethodPost, endpoint, param, utr)
	return
}

func (c *Client) PinFileFromFS(path string, name ...string) (pfr *PinFileResp, err error) {
	if len(name) > 1 {
		panic("name is only set one or ignored")
	}
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	pinName := fi.Name()
	if len(name) > 0 && len(name[0]) > 0 {
		pinName = name[0]
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	header := make([]byte, 512)
	_, err = file.Read(header)
	if err != nil && err != io.EOF {
		return nil, err
	}
	mime := http.DetectContentType(header)

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return c.PinFileFromStream(file, &model.UploadParam{
		Name: pinName,
		Size: fi.Size(),
		Mime: mime,
	})
}

func (c *Client) SetPinMeta(requestId string, meta *model.MetaData) (res *ResultBase, err error) {
	endpoint := "/pinning/meta/" + requestId
	res = new(ResultBase)
	err = c.processRequest(http.MethodPost, endpoint, meta, res)
	return
}

func (c *Client) SetPinPolicy(requestId string, param *model.PolicyParam) (res *ResultBase, err error) {
	endpoint := "/pinning/policy/" + requestId
	res = new(ResultBase)
	err = c.processRequest(http.MethodPost, endpoint, param, res)
	return
}

func (c *Client) SetUserPinPolicy(param *model.PolicyParam) (res *ResultBase, err error) {
	endpoint := "/pinning/policy"
	res = new(ResultBase)
	err = c.processRequest(http.MethodPost, endpoint, param, res)
	return
}
