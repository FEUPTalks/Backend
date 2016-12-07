package authentication

import (
	"fmt"
	"net/http"

	"github.com/FEUPTalks/Backend/util"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//RequireTokenAuthentication
func RequireTokenAuthentication(writer http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend, err := GetJWTAuthenticationBackend()
	if err != nil {
		util.ErrHandler(err, writer, http.StatusInternalServerError)
		return
	}

	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return authBackend.PublicKey, nil
		})

	//if err == nil && token.Valid && !authBackend.IsInBlacklist(req.Header.Get("Authorization")) {
	if err != nil {
		if err == request.ErrNoTokenInRequest {
			util.ErrHandler(err, writer, http.StatusUnauthorized)
		} else {
			util.ErrHandler(err, writer, http.StatusInternalServerError)
		}
		return
	}

	if token.Valid {
		next(writer, req)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}
