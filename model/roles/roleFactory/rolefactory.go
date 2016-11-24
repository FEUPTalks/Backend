package roleFactory

import "errors"
import "github.com/FEUPTalks/Backend/model/roles"

//GetRole used to create role implementations based on choice
func GetRole(role uint8) (roles.Role, error) {
	switch role {
	case 1:
		return &roles.AdminRole{}, nil
	case 2:
		return &roles.EmployeeRole{}, nil
	case 3:
		return &roles.UserRole{}, nil
	default:
		return nil, errors.New("Requested user role not available")
	}
}
