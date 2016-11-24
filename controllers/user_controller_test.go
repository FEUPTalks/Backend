package controllers_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"

	"github.com/FEUPTalks/Backend/controllers"
)

var (
	jsonStream = `{"userID": 1,"email": "teste@teste.com","name": "Teste","hashcode": "123456789abcdef","rolevalue": 3}`
)

func Example_UserController_CreateNewUser() {
	// userController := &controllers.UserController{}
	// request := httptest.NewRequest("POST", "/user", bytes.NewReader([]byte(jsonStream)))
	// writer := httptest.NewRecorder()
	// writer.Header().Set("Content-Type", "application/json")

	// userController.Create(writer, request)
	// fmt.Println(writer.Code)
	// userController.DeleteLastUser()
	// // Output:
	// //201

}
func Example_UserController_GetUserByID() {
	userController := &controllers.UserController{}
	request := httptest.NewRequest("GET", "/user", nil)
	writer := httptest.NewRecorder()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "1")

	request = request.WithContext(ctx)

	userController.GetUser(writer, request)

	fmt.Println(writer.Code)

	// Output:
	//200
}
func Example_UserController_GetUserByHashCode() {
	userController := &controllers.UserController{}
	request := httptest.NewRequest("GET", "/user", nil)
	writer := httptest.NewRecorder()

	ctx := context.Background()

	ctx = context.WithValue(ctx, "hashcode", "123456789abcdef")

	request = request.WithContext(ctx)

	userController.GetUser(writer, request)

	fmt.Println(writer.Code)

	// Output:
	//200
}
func Example_UserController_EditUser() {
	userController := &controllers.UserController{}
	request := httptest.NewRequest("POST", "/user", bytes.NewReader([]byte(jsonStream)))
	writer := httptest.NewRecorder()
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "1")
	request = request.WithContext(ctx)

	userController.SetUser(writer, request)
	fmt.Println(writer.Code)

	// Output:
	// 	200
}
