package model

//Profile
type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  uint8  `json:"role"`
	Token string `json:"token"`
}
