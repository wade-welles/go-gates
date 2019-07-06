package nand

import "testing"

func TestDMux(t *testing.T) {
	truthTable := []struct {
		in bool
		s  bool
		a  bool
		b  bool
	}{
		{true, true, false, true},
		{true, false, true, false},
		{false, false, false, false},
		{false, true, false, false},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.in, e.s), func(t *testing.T) {
			a, b := DMux(e.in, e.s)

			if (a != e.a) || (b != e.b) {
				t.Errorf("got %v and %v, want %v and %v", a, b, e.a, e.b)
			}
		})
	}
}
