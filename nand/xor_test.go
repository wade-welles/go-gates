package nand

import "testing"

func TestXor(t *testing.T) {
	truthTable := []struct {
		in  [2]bool
		out bool
	}{
		{[...]bool{true, true}, false},
		{[...]bool{true, false}, true},
		{[...]bool{false, false}, false},
		{[...]bool{false, true}, true},
	}

	for _, e := range truthTable {
		r := Xor(e.in[0], e.in[1])

		if r != e.out {
			t.Errorf("got %v, want %v", r, e.out)
		}
	}
}
