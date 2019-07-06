package nand

// And performs an AND operation on two inputs
func And(a bool, b bool) bool {
	return Exec(Exec(a, b), Exec(a, b))
}
