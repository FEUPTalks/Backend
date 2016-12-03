package roles

//UserRole role 3 of a user
type UserRole struct {
}

//Handle implementation Role interface
func (*UserRole) Handle() uint8 {
	return 3
}
