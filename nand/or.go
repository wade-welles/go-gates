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

// MultiWayOr performs an OR operation on all elements of the slice
func MultiWayOr(in []bool) bool {
	for _, e := range in {
		if e {
			return true
		}
	}

	return false
}
