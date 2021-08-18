package common

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewUTCTime(t *testing.T) {
	tm := NewUTCTime(time.Now())
	data, err := json.Marshal(tm)
	if err != nil {
		t.Fatal(err)
	}
	println(string(data))
}

func TestUnmarshalUTCTime(t *testing.T) {
	tm := UTCTime{}
	err := json.Unmarshal([]byte(`"2021-07-30T06:23:03Z"`), &tm)
	if err != nil {
		t.Fatal(err)
	}
	println(tm.String())
}
