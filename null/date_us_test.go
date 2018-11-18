package null

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCriarDateUSFrom(t *testing.T) {
	data1 := DateUSFrom("2015-01-17")
	data2 := NewDateUS(time.Date(2015, time.January, 17, 0, 0, 0, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.Equal(t, data1, data2)
	assert.True(t, data1.Valid)
}

func TestCriarDateUSFromInvalidaSeInputVazio(t *testing.T) {
	data1 := DateUSFrom("")
	data2 := NewDateUS(time.Date(2015, time.January, 17, 0, 0, 0, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.NotEqual(t, data1, data2)
	assert.False(t, data1.Valid)
}

func TestCriarDateUSNulo(t *testing.T) {
	data := DateUS{}

	var value *time.Time
	data.Scan(value)

	assert.False(t, data.Valid)
}

func TestLerDateUSNulo(t *testing.T) {
	data := DateUS{}

	var value *time.Time
	data.Scan(value)

	lido, _ := data.Value()

	assert.Nil(t, lido)
}

func TestCriarDateUS(t *testing.T) {
	data := DateUS{}

	data.Scan(time.Now())

	assert.True(t, data.Valid)
}

func TestConverterDateUSParaJson(t *testing.T) {
	data := DateUS{}

	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)
	data.Scan(agora)

	bytes, err := data.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, `"2015-04-17"`, string(bytes))
}

func TestConverterJsonParaDateUS(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)

	data := DateUS{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`"2015-04-17"`), &data)

	assert.Nil(t, err)
	assert.Equal(t, agora.String(), data.Date.String())
}

func TestConverterJsonParaDateUSNull(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)

	data := DateUS{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`null`), &data)

	assert.Nil(t, err)
	assert.False(t, data.Valid)
}

func TestCompararDateUS(t *testing.T) {
	data1 := NewDateUS(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))
	data2 := NewDateUS(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))

	assert.True(t, data1.Equal(&data2))
}
