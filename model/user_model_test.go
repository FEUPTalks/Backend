package model_test

import (
	"fmt"

	"github.com/FeupTalks/Backend/model"
	"github.com/FeupTalks/Backend/model/roles/roleFactory"
)

func Example_User_GetSetRole() {
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

	// Output:
	//3
	//1
	//2
	//3
}
