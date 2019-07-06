package nand

import (
	"reflect"
	"testing"
)

func TestMux(t *testing.T) {
	truthTable := []struct {
		in  [3]bool
		out bool
	}{
		{[...]bool{true, true, true}, true},
		{[...]bool{true, true, false}, true},
		{[...]bool{true, false, false}, true},
		{[...]bool{true, false, true}, false},
		{[...]bool{false, false, false}, false},
		{[...]bool{false, false, true}, false},
		{[...]bool{false, true, true}, true},
		{[...]bool{false, true, false}, false},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.in[:]), func(t *testing.T) {
			r := Mux(e.in[0], e.in[1], e.in[2])

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}

func TestMultiBitMux(t *testing.T) {
	truthTable := []struct {
		a   []bool
		b   []bool
		s   bool
		out []bool
	}{
		{
			a:   []bool{true, true, true},
			b:   []bool{false, true, false},
			s:   false,
			out: []bool{true, true, true},
		},
		{
			a:   []bool{false, true, false},
			b:   []bool{false, false, false},
			s:   true,
			out: []bool{false, false, false},
		},
	}

	for _, e := range truthTable {
		r := MultiBitMux(e.a, e.b, e.s)

		if !reflect.DeepEqual(r, e.out) {
			t.Errorf("got %v, want %v", r, e.out)
		}
	}
}
