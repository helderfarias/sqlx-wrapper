package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DDMMYYYYHHMMSS = "02/01/2006 15:04:05"
)

type DateTime struct {
	DateTime time.Time
	Valid    bool
}

func DateTimeFrom(src string) DateTime {
	dateTime, err := time.Parse(DDMMYYYYHHMMSS, src)
	if err != nil {
		return DateTime{DateTime: time.Now(), Valid: false}
	}

	return NewDateTime(dateTime)
}

func NewDateTime(datetime time.Time) DateTime {
	return DateTime{
		DateTime: time.Date(
			datetime.Year(),
			datetime.Month(),
			datetime.Day(),
			datetime.Hour(),
			datetime.Minute(),
			datetime.Second(),
			datetime.Nanosecond(),
			time.UTC),
		Valid: true,
	}
}

func (n *DateTime) Scan(value interface{}) error {
	n.DateTime, n.Valid = value.(time.Time)
	return nil
}

func (n DateTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.DateTime, nil
}

func (n DateTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	t := n.DateTime

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTime.MarshalJSON: year outside of range [0,9999]")
	}

	value := t.Format(DDMMYYYYHHMMSS)
	bytes := []byte(fmt.Sprintf("\"%s\"", value))
	return bytes, nil
}

func (n *DateTime) UnmarshalJSON(data []byte) (err error) {
	type dateDecode struct {
		DateTime string
		Valid    bool
	}

	if string(data) == "" || string(data) == "null" {
		n.Valid = false
		return nil
	}

	var decode dateDecode
	err = json.Unmarshal(data, &decode.DateTime)

	fmt.Println(decode.DateTime)

	n.DateTime, err = time.Parse(DDMMYYYYHHMMSS, string(decode.DateTime))
	n.Valid = decode.Valid
	n.Valid = err == nil
	return
}

func (n *DateTime) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	n.DateTime, err = time.Parse(DDMMYYYYHHMMSS, str)
	n.Valid = err == nil
	return err
}

func (n DateTime) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}

	t := n.DateTime

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTime.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(DDMMYYYYHHMMSS)), nil
}

func (d DateTime) IsZero() bool {
	return !d.Valid
}

func (d *DateTime) Equal(that *DateTime) bool {
	return d.DateTime.Equal(that.DateTime)
}

func (d *DateTime) Before(that *DateTime) bool {
	return d.DateTime.Before(that.DateTime)
}

func (d *DateTime) After(that *DateTime) bool {
	return d.DateTime.After(that.DateTime)
}
