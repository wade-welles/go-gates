package nand

import (
	"strconv"
	"testing"
)

var truthTable = []struct {
	in  [2]bool
	out bool
}{
	{[...]bool{true, true}, false},
	{[...]bool{true, false}, true},
	{[...]bool{false, false}, true},
	{[...]bool{false, true}, true},
}

func inToName(in [2]bool) string {
	first := strconv.FormatBool(in[0])
	second := strconv.FormatBool(in[1])

	return first + second
}

func TestExec(t *testing.T) {
	for _, e := range truthTable {
		t.Run(inToName(e.in), func(t *testing.T) {
			r := Exec(e.in[0], e.in[1])

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
