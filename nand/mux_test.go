package nand

import "testing"

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
