package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	DDMMYYYYHHMMSS_US = "2006-01-02 15:04:05"
)

type DateTimeUS struct {
	DateTimeUS time.Time
	Valid      bool
}

func DateTimeFromUS(src string) DateTimeUS {
	dateTime, err := time.Parse(DDMMYYYYHHMMSS_US, src)
	if err != nil {
		return DateTimeUS{DateTimeUS: time.Now(), Valid: false}
	}

	return NewDateTimeUS(dateTime)
}

func NewDateTimeUS(datetime time.Time) DateTimeUS {
	return DateTimeUS{
		DateTimeUS: time.Date(
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

func (n *DateTimeUS) Scan(value interface{}) error {
	n.DateTimeUS, n.Valid = value.(time.Time)
	return nil
}

func (n DateTimeUS) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.DateTimeUS, nil
}

func (n DateTimeUS) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	t := n.DateTimeUS

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTimeUS.MarshalJSON: year outside of range [0,9999]")
	}

	value := t.Format(DDMMYYYYHHMMSS_US)
	bytes := []byte(fmt.Sprintf("\"%s\"", value))
	return bytes, nil
}

func (n *DateTimeUS) UnmarshalJSON(data []byte) (err error) {
	type dateDecode struct {
		DateTimeUS string
		Valid      bool
	}

	if string(data) == "" || string(data) == "null" {
		n.Valid = false
		return nil
	}

	var decode dateDecode
	err = json.Unmarshal(data, &decode.DateTimeUS)

	fmt.Println(decode.DateTimeUS)

	n.DateTimeUS, err = time.Parse(DDMMYYYYHHMMSS_US, string(decode.DateTimeUS))
	n.Valid = decode.Valid
	n.Valid = err == nil
	return
}

func (n *DateTimeUS) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	n.DateTimeUS, err = time.Parse(DDMMYYYYHHMMSS_US, str)
	n.Valid = err == nil
	return err
}

func (n DateTimeUS) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}

	t := n.DateTimeUS

	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("DateTimeUS.MarshalJSON: year outside of range [0,9999]")
	}

	return []byte(t.Format(DDMMYYYYHHMMSS_US)), nil
}

func (d DateTimeUS) IsZero() bool {
	return !d.Valid
}

func (d *DateTimeUS) Equal(that *DateTimeUS) bool {
	return d.DateTimeUS.Equal(that.DateTimeUS)
}

func (d *DateTimeUS) Before(that *DateTimeUS) bool {
	return d.DateTimeUS.Before(that.DateTimeUS)
}

func (d *DateTimeUS) After(that *DateTimeUS) bool {
	return d.DateTimeUS.After(that.DateTimeUS)
}

func (d *DateTimeUS) String() string {
	bytes, err := d.MarshalText()
	if err != nil {
		return ""
	}

	return string(bytes)
}
