package null

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCriarDateTimeNulo(t *testing.T) {
	data := DateTime{}

	var value *time.Time
	data.Scan(value)

	assert.False(t, data.Valid)
}

func TestLerDateTimeNulo(t *testing.T) {
	data := DateTime{}

	var value *time.Time
	data.Scan(value)

	lido, _ := data.Value()

	assert.Nil(t, lido)
}

func TestCriarDateTime(t *testing.T) {
	data := DateTime{}

	data.Scan(time.Now())

	assert.True(t, data.Valid)
}

func TestConverterDateTimeParaJson(t *testing.T) {
	data := DateTime{}

	agora := time.Date(2015, time.April, 17, 8, 10, 10, 10, time.UTC)
	data.Scan(agora)

	bytes, err := data.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, `"17/04/2015 08:10:10"`, string(bytes))
}

func TestConverterJsonParaDateTime(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 8, 10, 10, 0, time.UTC)

	data := DateTime{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`"17/04/2015 08:10:10"`), &data)

	assert.Nil(t, err)
	assert.Equal(t, agora.String(), data.DateTime.String())
}

func TestCompararDateTime(t *testing.T) {
	data1 := NewDateTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))
	data2 := NewDateTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))

	assert.True(t, data1.Equal(&data2))
}
