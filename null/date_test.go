package null

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCriarDateNulo(t *testing.T) {
	data := Date{}

	var value *time.Time
	data.Scan(value)

	assert.False(t, data.Valid)
}

func TestLerDateNulo(t *testing.T) {
	data := Date{}

	var value *time.Time
	data.Scan(value)

	lido, _ := data.Value()

	assert.Nil(t, lido)
}

func TestCriarDate(t *testing.T) {
	data := Date{}

	data.Scan(time.Now())

	assert.True(t, data.Valid)
}

func TestConverterDateParaJson(t *testing.T) {
	data := Date{}

	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)
	data.Scan(agora)

	bytes, err := data.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, `"17/04/2015"`, string(bytes))
}

func TestConverterJsonParaDate(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)

	data := Date{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`"17/04/2015"`), &data)

	assert.Nil(t, err)
	assert.Equal(t, agora.String(), data.Date.String())
}

func TestConverterJsonParaDateNull(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)

	data := Date{}
	data.Scan(agora)

	err := json.Unmarshal([]byte(`null`), &data)

	assert.Nil(t, err)
	assert.False(t, data.Valid)
}

func TestCompararDate(t *testing.T) {
	data1 := NewDate(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))
	data2 := NewDate(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))

	assert.True(t, data1.Equal(&data2))
}
