package nand

import "testing"

func TestAnd(t *testing.T) {
	truthTable := []struct {
		in  [2]bool
		out bool
	}{
		{[...]bool{true, true}, true},
		{[...]bool{true, false}, false},
		{[...]bool{false, false}, false},
		{[...]bool{false, true}, false},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.in), func(t *testing.T) {
			r := And(e.in[0], e.in[1])

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
