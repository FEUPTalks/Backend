package parameters

//TokenAuthentication token dto
type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}
