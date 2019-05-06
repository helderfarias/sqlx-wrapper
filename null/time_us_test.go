package null

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTimeString(t *testing.T) {
	time1 := TimeFrom("12:56:22")
	time2 := NewTime(time.Date(2015, time.January, 17, 10, 20, 05, 5, time.UTC))
	time3 := TimeFrom("")
	time4 := TimeFrom("12:56")

	assert.Equal(t, "12:56:22", time1.String())
	assert.Equal(t, "10:20:05", time2.String())
	assert.Empty(t, time3.String())
	assert.Equal(t, "12:56:00", time4.String())
}

func TestCriarTimeFrom(t *testing.T) {
	data1 := TimeFrom("10:10:01")
	data2 := NewTime(time.Date(2015, time.January, 17, 10, 10, 1, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.Equal(t, data1, data2)
	assert.True(t, data1.Valid)
}

func TestCriarTimeFromInvalidaSeInputVazio(t *testing.T) {
	data1 := TimeFrom("")
	data2 := NewTime(time.Date(2015, time.January, 17, 0, 0, 0, 0, time.UTC))

	assert.NotNil(t, data1)
	assert.NotEqual(t, data1, data2)
	assert.False(t, data1.Valid)
}

func TestCriarTimeNulo(t *testing.T) {
	t1 := Time{}

	var value *time.Time
	t1.Scan(value)

	assert.False(t, t1.Valid)
}

func TestLerTimeNulo(t *testing.T) {
	t1 := Time{}

	var value *time.Time
	t1.Scan(value)

	lido, _ := t1.Value()

	assert.Nil(t, lido)
}

func TestCriarTime(t *testing.T) {
	t1 := Time{}

	t1.Scan(time.Now())

	assert.True(t, t1.Valid)
}

func TestConverterTimeParaJson(t *testing.T) {
	t1 := Time{}

	agora := time.Date(2015, time.April, 17, 10, 10, 10, 0, time.UTC)
	t1.Scan(agora)

	bytes, err := t1.MarshalJSON()

	assert.Nil(t, err)
	assert.Equal(t, `"10:10:10"`, string(bytes))
}

func TestConverterJsonParaTime(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 10, 10, 10, 0, time.UTC)

	t1 := Time{}
	t1.Scan(agora)

	err := json.Unmarshal([]byte(`"10:10"`), &t1)

	assert.Nil(t, err)
	assert.Equal(t, "10:10:00", t1.String())
}

func TestConverterJsonParaTimeNull(t *testing.T) {
	agora := time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC)

	t1 := Time{}
	t1.Scan(agora)

	err := json.Unmarshal([]byte(`null`), &t1)

	assert.Nil(t, err)
	assert.False(t, t1.Valid)
}

func TestCompararTime(t *testing.T) {
	t1 := NewTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))
	t2 := NewTime(time.Date(2015, time.April, 17, 0, 0, 0, 0, time.UTC))

	assert.True(t, t1.Equal(&t2))
}
