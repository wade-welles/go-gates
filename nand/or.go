package nand

// Or performs an OR operation on two inputs
func Or(a bool, b bool) bool {
	return Exec(Not(a), Not(b))
}
