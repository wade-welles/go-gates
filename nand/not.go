package nand

// Not performs a NOT operation on the input.
func Not(in bool) bool {
	return Exec(in, in)
}

// MultiBitNot performs a NOT operation on each element of the slice
func MultiBitNot(in []bool) []bool {
	r := make([]bool, len(in))

	for i, e := range in {
		r[i] = Not(e)
	}

	return r
}
