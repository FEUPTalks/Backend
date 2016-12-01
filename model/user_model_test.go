package test

import (
	"fmt"

	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/model/roles/roleFactory"
	"testing"
)

/*
Expect output:
3
1
2
3
 */
func TestGetSetRole(t *testing.T) {
	roleTest1, _ := roleFactory.GetRole(1)
	roleTest2, _ := roleFactory.GetRole(2)
	roleTest3, _ := roleFactory.GetRole(3)

	exampleUser := model.NewUser()
	fmt.Println(exampleUser.GetRoleValue())

	exampleUser.SetRole(roleTest1)
	fmt.Println(exampleUser.GetRoleValue())

	exampleUser.SetRole(roleTest2)
	fmt.Println(exampleUser.GetRoleValue())

	exampleUser.SetRole(roleTest3)
	fmt.Println(exampleUser.GetRoleValue())
}
