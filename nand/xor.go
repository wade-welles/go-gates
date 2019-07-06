package nand

// Xor performs an XOR operation on two inputs
func Xor(a, b bool) bool {
	return Or(And(a, Not(b)), And(b, Not(a)))
}
