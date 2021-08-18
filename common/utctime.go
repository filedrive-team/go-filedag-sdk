package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UTCTime struct {
	time.Time
}

func NewUTCTime(t time.Time) UTCTime {
	nt := UTCTime{}
	nt.Time = t.UTC()
	return nt
}

func (t UTCTime) MarshalJSON() ([]byte, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return []byte(`"-"`), nil
	}
	tune := t.UTC().Format(fmt.Sprintf(`"%s"`, time.RFC3339))
	return []byte(tune), nil
}

func (t *UTCTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `"-"` {
		return nil
	}

	var err error
	t.Time, err = time.Parse(fmt.Sprintf(`"%s"`, time.RFC3339), string(data))
	return err
}

func (t UTCTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *UTCTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = UTCTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *UTCTime) String() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.UTC().Format(time.RFC3339)
}

func (t *UTCTime) YearMonthDate() string {
	if t.IsEmpty() {
		return ""
	}
	return t.Time.UTC().Format("2006-01-02")
}

func (t *UTCTime) IsEmpty() bool {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return true
	}
	return false
}
