package nand

import (
	"reflect"
	"testing"
)

func TestOr(t *testing.T) {
	truthTable := []struct {
		in  [2]bool
		out bool
	}{
		{[...]bool{true, true}, true},
		{[...]bool{true, false}, true},
		{[...]bool{false, false}, false},
		{[...]bool{false, true}, true},
	}

	for _, e := range truthTable {
		r := Or(e.in[0], e.in[1])

		if r != e.out {
			t.Errorf("got %v, want %v", r, e.out)
		}
	}
}

func TestMultiBitOr(t *testing.T) {
	a := []bool{true, false, true, false}
	b := []bool{false, false, true, true}
	exp := []bool{true, false, true, true}
	r := MultiBitOr(a, b)

	if !reflect.DeepEqual(r, exp) {
		t.Errorf("got %v, want %v", r, exp)
	}
}
