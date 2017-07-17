package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DDMMYYYY = "02/01/2006"
)

type Date struct {
	Date  time.Time
	Valid bool
}

func NewDate(date time.Time) Date {
	return Date{
		Date:  time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
		Valid: true,
	}
}

func (n *Date) Scan(value interface{}) error {
	n.Date, n.Valid = value.(time.Time)
	return nil
}

func (n Date) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date, nil
}

func (n Date) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	t := n.Date

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	value := t.Format(DDMMYYYY)
	bytes := []byte(fmt.Sprintf("\"%s\"", value))
	return bytes, nil
}

func (n *Date) UnmarshalJSON(data []byte) (err error) {
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

	n.Date, err = time.Parse(DDMMYYYY, string(decode.Date))
	n.Valid = decode.Valid
	n.Valid = err == nil
	return
}

func (n *Date) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	n.Date, err = time.Parse(DDMMYYYY, str)
	n.Valid = err == nil
	return err
}

func (n Date) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}

	t := n.Date

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(DDMMYYYY)), nil
}

func (d Date) IsZero() bool {
	return !d.Valid
}

func (d *Date) Equal(that *Date) bool {
	return d.Date.Equal(that.Date)
}

func (d *Date) Before(that *Date) bool {
	return d.Date.Before(that.Date)
}

func (d *Date) After(that *Date) bool {
	return d.Date.After(that.Date)
}
