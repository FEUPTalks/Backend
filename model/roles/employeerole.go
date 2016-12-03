package roles

//EmployeeRole role 2 of a user
type EmployeeRole struct {
}

//Handle implementation Role interface
func (*EmployeeRole) Handle() uint8 {
	return 2
}
