// Package nand implements the NAND gate and other gates based upon it.
package nand

// Exec performs a NAND operation on two inputs.
func Exec(a, b bool) bool {
	return !(a && b)
}
