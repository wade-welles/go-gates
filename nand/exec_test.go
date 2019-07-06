package nand

import (
	"strconv"
	"testing"
)

func inToName(in ...bool) string {
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
		a   bool
		b   bool
		out bool
	}{
		{true, true, false},
		{true, false, true},
		{false, false, true},
		{false, true, true},
	}

	for _, e := range truthTable {
		t.Run(inToName(e.a, e.b), func(t *testing.T) {
			r := Exec(e.a, e.b)

			if r != e.out {
				t.Errorf("got %v, want %v", r, e.out)
			}
		})
	}
}
