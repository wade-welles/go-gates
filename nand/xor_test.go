package nand

import "testing"

func TestXor(t *testing.T) {
	truthTable := []struct {
		a   bool
		b   bool
		out bool
	}{
		{true, true, false},
		{true, false, true},
		{false, false, false},
		{false, true, true},
	}

	for _, e := range truthTable {
		r := Xor(e.a, e.b)

		if r != e.out {
			t.Errorf("got %v, want %v", r, e.out)
		}
	}
}
