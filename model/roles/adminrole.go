package roles

//AdminRole role 1 of a user
type AdminRole struct {
}

//Handle implementation role interface
func (*AdminRole) Handle() uint8 {
	return 1
}
