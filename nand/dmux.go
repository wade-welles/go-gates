package nand

// DMux outputs the provided input in a or b based on s
func DMux(in, s bool) (a, b bool) {
	a = And(Not(s), in)
	b = And(s, in)

	return
}
