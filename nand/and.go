package nand

// And performs an AND operation on two inputs
func And(a bool, b bool) bool {
	return Exec(Exec(a, b), Exec(a, b))
}

// MultiBitAnd performs an AND operation on each element of the slice
func MultiBitAnd(a []bool, b []bool) []bool {
	r := make([]bool, len(a))

	for i := range r {
		r[i] = And(a[i], b[i])
	}

	return r
}
