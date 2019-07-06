package nand

// Exec performs a NAND operation on two inputs
func Exec(a bool, b bool) bool {
	return !(a && b)
}
