package nand

// Mux outputs one of two inputs based on s
func Mux(a bool, b bool, s bool) bool {
	return Or(
		And(a, Not(s)),
		And(b, s),
	)
}
