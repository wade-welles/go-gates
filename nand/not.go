package nand

// Not performs a NOT operation on the input.
func Not(a bool) bool {
	return Exec(a, a)
}
