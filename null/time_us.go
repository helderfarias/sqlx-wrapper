package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const (
	TIME_FORMAT_HHMMSS = "15:04:05"
	TIME_FORMAT_HHMM   = "15:04"
)

type Time struct {
	Date  time.Time
	Valid bool
}

func TimeFrom(src string) Time {
	date, err := parseTimeWithFallback(src)
	if err != nil {
		fallback := time.Now()
		return Time{
			Date:  time.Date(0, 0, 0, fallback.Hour(), fallback.Minute(), fallback.Second(), fallback.Nanosecond(), time.UTC),
			Valid: false,
		}
	}

	return NewTime(date)
}

func NewTime(date time.Time) Time {
	return Time{
		Date:  time.Date(0, 0, 0, date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), time.UTC),
		Valid: true,
	}
}

func (n *Time) Scan(value interface{}) error {
	n.Date, n.Valid = value.(time.Time)
	return nil
}

func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date, nil
}

func (n Time) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	t := n.Date

	value := t.Format(TIME_FORMAT_HHMMSS)
	bytes := []byte(fmt.Sprintf("\"%s\"", value))
	return bytes, nil
}

func (n *Time) UnmarshalJSON(data []byte) (err error) {
	type dateDecode struct {
		Date  string
		Valid bool
	}

	if string(data) == "" || string(data) == "null" {
		n.Date = time.Time{}
		n.Valid = false
		err = nil
		return
	}

	var decode dateDecode
	err = json.Unmarshal(data, &decode.Date)

	n.Date, err = parseTimeWithFallback(string(decode.Date))
	n.Valid = decode.Valid
	n.Valid = err == nil
	return
}

func (n *Time) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	n.Date, err = parseTimeWithFallback(str)
	n.Valid = err == nil
	return err
}

func (n Time) MarshalText() ([]byte, error) {
	if !n.Valid {
		return []byte{}, nil
	}

	t := n.Date

	return []byte(t.Format(TIME_FORMAT_HHMMSS)), nil
}

func (d Time) IsZero() bool {
	return !d.Valid
}

func (d *Time) Equal(that *Time) bool {
	return d.Date.Hour() == that.Date.Hour() &&
		d.Date.Minute() == that.Date.Minute() &&
		d.Date.Second() == that.Date.Second()
}

func (d *Time) String() string {
	bytes, err := d.MarshalText()
	if err != nil {
		return ""
	}

	return string(bytes)
}

func parseTimeWithFallback(src string) (time.Time, error) {
	if date, err := time.Parse(TIME_FORMAT_HHMMSS, src); err == nil {
		return date, nil
	}

	return time.Parse(TIME_FORMAT_HHMM, src)
}
