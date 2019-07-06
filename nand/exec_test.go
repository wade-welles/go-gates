package nand

import (
	"strconv"
	"testing"
)

func inToName(in []bool) string {
	var res string

	for i, b := range in {
		if i == 0 {
			res += strconv.FormatBool(b)
		} else {
			res += "_" + strconv.FormatBool(b)
		}
	}

	return res
}

func TestExec(t *testing.T) {
	truthTable := []struct {
		in  [2]bool
		out bool
	}{
		{[...]bool{true, true}, false},
		{[...]bool{true, false}, true},
		{[...]bool{false, false}, true},
		{[...]bool{false, true}, true},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.in[:]), func(t *testing.T) {
			r := Exec(e.in[0], e.in[1])

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
