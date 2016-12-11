package services

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/FEUPTalks/Backend/core/authentication"
	"github.com/FEUPTalks/Backend/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//Login user login function
func Login(requestUser *model.LoginInfo) (*model.Profile, error) {
	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user, err := authBackend.Authenticate(requestUser)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token, err := authBackend.GenerateToken(user.UUID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.Profile{user.Name, user.Email, user.RoleValue, token}, nil
}

//RefreshToken
func RefreshToken(requestUser *model.User) ([]byte, error) {
	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response, err := json.Marshal(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}

func Logout(req *http.Request) error {
	authBackend, err := authentication.GetJWTAuthenticationBackend()
	if err != nil {
		log.Println(err)
		return err
	}

	tokenRequest, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
