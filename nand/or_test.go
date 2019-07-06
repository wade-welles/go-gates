package nand

import (
	"reflect"
	"testing"
)

func TestOr(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		out bool
	}{
		{true, true, true},
		{true, false, true},
		{false, false, false},
		{false, true, true},
	}

	for _, e := range truthTable {
		r := Or(e.a, e.b)

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

func TestMultiWayOr(t *testing.T) {
	truthTable := []struct {
		name string
		in   []bool
		out  bool
	}{
		{"all inputs false", []bool{false, false, false}, false},
		{"at least one input true", []bool{true, false, false}, true},
	}

	for _, e := range truthTable {
		t.Run(e.name, func(t *testing.T) {
			r := MultiWayOr(e.in)

			if e.out != r {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
