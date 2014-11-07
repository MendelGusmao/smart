package smart

import (
	"reflect"
	"testing"
)

func TestSmart(t *testing.T) {
	d := ","
	e := "\\"

	tests := map[string][]string{
		"a,b,c":     []string{"a", "b", "c"},
		"a,b,c\\,d": []string{"a", "b", "c,d"},
	}

	for str, exp := range tests {
		slc := Split(str, d, e)

		if !reflect.DeepEqual(slc, exp) {
			t.Fatalf("testing %s failed: expected %v, got %v", str, exp, slc)
		}
	}
}
