package null

import (
	"testing"
)

var (
	intJSON         = []byte(`12345`)
	nullIntJSON     = []byte(`{"Int64":12345,"Valid":true}`)
	boolJSON        = []byte(`true`)
	falseJSON       = []byte(`false`)
	nullBoolJSON    = []byte(`{"Bool":true,"Valid":true}`)
	floatJSON       = []byte(`1.2345`)
	nullFloatJSON   = []byte(`{"Float64":1.2345,"Valid":true}`)
	stringJSON      = []byte(`"test"`)
	blankStringJSON = []byte(`""`)
	nullStringJSON  = []byte(`{"String":"test","Valid":true}`)
	nullJSON        = []byte(`null`)
)

func assertJSONEquals(t *testing.T, data []byte, cmp string, from string) {
	if string(data) != cmp {
		t.Errorf("bad %s data: %s ≠ %s\n", from, data, cmp)
	}
}

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}
