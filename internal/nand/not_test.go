package nand

import (
	"strconv"
	"testing"
)

func TestNot(t *testing.T) {
	var truthTable = []struct {
		in  bool
		out bool
	}{
		{true, false},
		{false, true},
	}

	for _, e := range truthTable {
		t.Run(strconv.FormatBool(e.in), func(t *testing.T) {
			r := Not(e.in)

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
