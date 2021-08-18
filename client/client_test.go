package client

import (
	"bytes"
	"github.com/filedrive-team/go-filedag-sdk/common"
	"github.com/filedrive-team/go-filedag-sdk/model"
	"net/http"
	"os"
	"path"
	"testing"
	"time"
)

const (
	host = "https://api.filedag.cloud"

	apiKey    = "YOUR_API_KEY"
	apiSecret = "YOUR_API_SECRET"
	// or
	jwt = "YOUR_JWT"
)

func TestClient_ListPin(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	before := common.NewUTCTime(time.Now())
	after := common.NewUTCTime(time.Now())
	params := []*model.SearchPinParam{
		{
			Status: []model.Status{model.PINNED},
			Limit:  1,
		},
		{
			Cid:    []string{"QmYcpdUffAwufMSLPjH8CzjAHEqxmF6NdTrqJWvjw9LbUT"},
			Status: []model.Status{model.PINNED},
			Limit:  1,
		},
		{
			Name:   "aa",
			Match:  model.PARTIAL,
			Status: []model.Status{model.PINNED},
			Limit:  1,
		},
		{
			Name:   "aa",
			Match:  model.PARTIAL,
			Status: []model.Status{model.PINNED, model.PINNING},
			Limit:  1,
		},
		{
			Status: []model.Status{model.PINNED, model.PINNING},
			Before: &before,
			Limit:  1,
		},
		{
			Status: []model.Status{model.PINNED, model.PINNING},
			After:  &after,
			Limit:  1,
		},
		{
			Status: []model.Status{model.PINNED, model.PINNING},
			Meta: map[string]string{
				"app_id": "app_id",
			},
			Limit: 1,
		},
	}
	for _, param := range params {
		resp, err := cli.ListPin(param)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Code != http.StatusOK {
			t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
		} else {
			t.Logf("code=%v data=%+v", resp.Code, resp.Data)
		}
	}
}

func TestClient_AddPin(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	pin := &model.Pin{
		Cid:  "QmZ9Uh1qrkDuFzRxDYPZibLqd7c9nQ5b3DjkTzVaKpzrJ6",
		Name: "test1",
		Meta: map[string]string{"id": "1"},
	}
	resp, err := cli.AddPin(pin)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusAccepted {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_GetPin(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.GetPin("31bea2b4-899e-40e7-8f4d-ac7c93e66697")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_ReplacePin(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	pin := &model.Pin{
		Cid:  "QmZ9Uh1qrkDuFzRxDYPZibLqd7c9nQ5b3DjkTzVaKpzrJ6",
		Name: "test2",
		Meta: map[string]string{"id": "2"},
	}
	resp, err := cli.ReplacePin("7fa89c98-507d-41c8-85c4-005c84cc9ae3", pin)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusAccepted {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_RemovePin(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.RemovePin("31bea2b4-899e-40e7-8f4d-ac7c93e66697")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusAccepted {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v", resp.Code)
	}
}

func TestClient_GenerateApiKey(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	admin := true
	resp, err := cli.GenerateApiKey(&model.KeyInfo{
		KeyName: "testKey",
		Scope: model.Scope{
			Admin: &admin,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_RevokeApiKey(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.RevokeApiKey("8f17fcc078f96cda9837")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v", resp.Code)
	}
}

func TestClient_PinnedDataTotal(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.PinnedDataTotal()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_SetPinMeta(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.SetPinMeta("31bea2b4-899e-40e7-8f4d-ac7c93e66697", &model.MetaData{
		Name: "test_2",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v", resp.Code)
	}
}

func TestClient_SetPinPolicy(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.SetPinPolicy("31bea2b4-899e-40e7-8f4d-ac7c93e66697", &model.PolicyParam{
		NewPinPolicy: map[string]int{
			"wulan-01": 1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusAccepted {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v", resp.Code)
	}
}

func TestClient_SetUserPinPolicy(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	resp, err := cli.SetUserPinPolicy(&model.PolicyParam{
		NewPinPolicy: map[string]int{
			"shenzhen-01": 1,
			"wulan-01":    1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v", resp.Code)
	}
}

func generateTestFile(t *testing.T) string {
	path := path.Join(os.TempDir(), "filedag.txt")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	file.WriteString("Hello, FileDAG!")
	return path
}

func TestClient_PinFileFromFS(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	path := generateTestFile(t)
	resp, err := cli.PinFileFromFS(path, "filedag.txt")
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_PinFileFromStream(t *testing.T) {
	cli := NewWithJwtToken(host, jwt)
	buf := bytes.NewBufferString("Hello, FileDAG!")
	mime := http.DetectContentType(buf.Bytes())
	resp, err := cli.PinFileFromStream(buf, &model.UploadParam{
		Name: "filedag",
		Mime: mime,
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}

func TestClient_PinnedDataTotalWithKeySecret(t *testing.T) {
	cli := NewWithKeySecret(host, apiKey, apiSecret)
	resp, err := cli.PinnedDataTotal()
	if err != nil {
		t.Fatal(err)
	}
	if resp.Code != http.StatusOK {
		t.Errorf("code=%v failure=%+v", resp.Code, resp.Failure)
	} else {
		t.Logf("code=%v data=%+v", resp.Code, resp.Data)
	}
}
