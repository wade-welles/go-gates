package nand

// Mux outputs one of two inputs based on s
func Mux(a bool, b bool, s bool) bool {
	return Or(
		And(a, Not(s)),
		And(b, s),
	)
}

// MultiBitMux outputs, based on s, a slice containing either the elements of a or b
func MultiBitMux(a []bool, b []bool, s bool) []bool {
	r := make([]bool, len(a))

	for i := range r {
		r[i] = Mux(a[i], b[i], s)
	}

	return r
}
