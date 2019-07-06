package nand

// Mux outputs one of two inputs based on the s value
func Mux(a bool, b bool, s bool) bool {
	return Or(
		And(a, Not(s)),
		And(b, s),
	)
}
