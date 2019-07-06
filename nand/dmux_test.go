package nand

import "testing"

func TestDMux(t *testing.T) {
	truthTable := []struct {
		in  [2]bool
		out [2]bool
	}{
		{[...]bool{true, true}, [...]bool{false, true}},
		{[...]bool{true, false}, [...]bool{true, false}},
		{[...]bool{false, false}, [...]bool{false, false}},
		{[...]bool{false, true}, [...]bool{false, false}},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.in[:]), func(t *testing.T) {
			a, b := DMux(e.in[0], e.in[1])

			if (a != e.out[0]) || (b != e.out[1]) {
				t.Errorf("got %v and %v, want %v and %v", a, b, e.out[0], e.out[1])
			}
		})
	}
}
