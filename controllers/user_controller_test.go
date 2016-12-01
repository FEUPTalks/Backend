package controllers

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
)

var (
	jsonStream = `{"userID": 1,"email": "em07152@fe.up.pt","name": "Teste","hashcode": "123456789abcdef","rolevalue": 3}`
)

/*
Expect Http Code: 201
 */
func TestCreateNewUser(t *testing.T) {
	userController := &UserController{}
	request := httptest.NewRequest("POST", "/user", bytes.NewReader([]byte(jsonStream)))
	writer := httptest.NewRecorder()
	writer.Header().Set("Content-Type", "application/json")

	userController.Create(writer, request)
	fmt.Println(writer.Code)
	userController.DeleteLastUser()
}

/*
Expect Http Code: 200
 */
func TestGetUserByID(t *testing.T) {
	userController := &UserController{}
	request := httptest.NewRequest("GET", "/user", nil)
	writer := httptest.NewRecorder()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "1")
	request = request.WithContext(ctx)

	userController.GetUser(writer, request)
	fmt.Println(writer.Code)
}

/*
Expect Http Code: 200
 */
func TestGetUserByHashCode(t *testing.T) {
	userController := &UserController{}
	request := httptest.NewRequest("GET", "/user", nil)
	writer := httptest.NewRecorder()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "hashcode", "123456789abcdef")
	request = request.WithContext(ctx)

	userController.GetUser(writer, request)
	fmt.Println(writer.Code)
}

/*
Expect Http Code: 200
 */
func TestEditUser(t *testing.T) {
	userController := &UserController{}
	request := httptest.NewRequest("PUT", "/user", bytes.NewReader([]byte(jsonStream)))
	writer := httptest.NewRecorder()
	writer.Header().Set("Content-Type", "application/json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "1")
	request = request.WithContext(ctx)

	userController.SetUser(writer, request)
	fmt.Println(writer.Code)
}
