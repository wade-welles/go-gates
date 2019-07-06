package nand

// Or performs an OR operation on two inputs
func Or(a bool, b bool) bool {
	return Exec(Not(a), Not(b))
}

// MultiBitOr performs an OR operation on each element of the slice
func MultiBitOr(a []bool, b []bool) []bool {
	r := make([]bool, len(a))

	for i := range r {
		r[i] = Or(a[i], b[i])
	}

	return r
}
