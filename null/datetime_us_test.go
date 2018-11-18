package null

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCriarDateTimeUSFrom(t *testing.T) {
	data1 := DateTimeFromUS("2015-04-17 08:10:10")
	data2 := NewDateTimeUS(time.Date(2015, time.April, 17, 8, 10, 10, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.Equal(t, data1, data2)
	assert.True(t, data1.Valid)
}

func TestCriarDateTimeUSFromInvalidaSeInputVazio(t *testing.T) {
	data1 := DateTimeFromUS("")
	data2 := NewDateTimeUS(time.Date(2015, time.January, 17, 0, 0, 0, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.NotEqual(t, data1, data2)
	assert.False(t, data1.Valid)
}

func TestCriarDateTimeUSNulo(t *testing.T) {
	data := DateTimeUS{}

	var value *time.Time
	data.Scan(value)

	assert.False(t, data.Valid)
}

func TestLerDateTimeUSNulo(t *testing.T) {
	data := DateTimeUS{}

	var value *time.Time
	data.Scan(value)

	lido, _ := data.Value()

	assert.Nil(t, lido)
}

func TestCriarDateTimeUS(t *testing.T) {
	data := DateTimeUS{}

	data.Scan(time.Now())

	assert.True(t, data.Valid)
}

func TestConverterDateTimeUSParaJson(t *testing.T) {
	data := DateTimeUS{}

	agora := time.Date(2015, time.April, 17, 8, 10, 10, 10, time.UTC)
	data.Scan(agora)

	bytes, err := data.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, `"2015-04-17 08:10:10"`, string(bytes))
}

func TestConverterJsonParaDateTimeUS(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 8, 10, 10, 0, time.UTC)

	data := DateTimeUS{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`"2015-04-17 08:10:10"`), &data)

	assert.Nil(t, err)
	assert.Equal(t, agora.String(), data.DateTimeUS.String())
}

func TestCompararDateTimeUS(t *testing.T) {
	data1 := NewDateTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))
	data2 := NewDateTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))

	assert.True(t, data1.Equal(&data2))
}
