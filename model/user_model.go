package model

import (
	"github.com/FEUPTalks/Backend/model/roles"
	"github.com/FEUPTalks/Backend/model/roles/roleFactory"
)

//User struct
type User struct {
	UserID    int    `json:"userID"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	HashCode  string `json:"hashcode"`
	RoleValue uint8  `json:"rolevalue"`
	role      roles.Role
}

//NewUser creates an empty user with a role user
func NewUser() *User {
	userRole, _ := roleFactory.GetRole(3)
	user := &User{}
	user.SetRole(userRole)
	return user
}

//GetRoleValue returns value of the role the user is in
func (user *User) GetRoleValue() uint8 {
	return user.role.Handle()
}

//SetRole changs the role of the user
func (user *User) SetRole(role roles.Role) {
	user.role = role
	user.RoleValue = user.role.Handle()
}
