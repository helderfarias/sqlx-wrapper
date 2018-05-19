package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DDMMYYYY_US = "2006-01-02"
)

type DateUS struct {
	Date  time.Time
	Valid bool
}

func NewDateUS(date time.Time) DateUS {
	return DateUS{
		Date:  time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
		Valid: true,
	}
}

func (n *DateUS) Scan(value interface{}) error {
	n.Date, n.Valid = value.(time.Time)
	return nil
}

func (n DateUS) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date, nil
}

func (n DateUS) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte(""), nil
	}

	t := n.Date

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	value := t.Format(DDMMYYYY_US)
	bytes := []byte(fmt.Sprintf("\"%s\"", value))
	return bytes, nil
}

func (n *DateUS) UnmarshalJSON(data []byte) (err error) {
	type dateDecode struct {
		Date  string
		Valid bool
	}

	if string(data) == "null" {
		n.Date = time.Time{}
		n.Valid = false
		err = nil
		return
	}

	var decode dateDecode
	err = json.Unmarshal(data, &decode.Date)

	n.Date, err = time.Parse(DDMMYYYY_US, string(decode.Date))
	n.Valid = decode.Valid
	n.Valid = err == nil
	return
}

func (n *DateUS) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	n.Date, err = time.Parse(DDMMYYYY_US, str)
	n.Valid = err == nil
	return err
}

func (n DateUS) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}

	t := n.Date

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(DDMMYYYY_US)), nil
}

func (d DateUS) IsZero() bool {
	return !d.Valid
}

func (d *DateUS) Equal(that *Date) bool {
	return d.Date.Equal(that.Date)
}

func (d *DateUS) Before(that *Date) bool {
	return d.Date.Before(that.Date)
}

func (d *DateUS) After(that *Date) bool {
	return d.Date.After(that.Date)
}
