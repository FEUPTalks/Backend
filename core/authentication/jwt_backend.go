package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"

	"log"

	"database/sql"

	"github.com/FEUPTalks/Backend/database"
	"github.com/FEUPTalks/Backend/model"
	"github.com/FEUPTalks/Backend/settings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type jwtAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset  = 3600
)

var authBackendInstance *jwtAuthenticationBackend
var once sync.Once

//ErrUserNotFound used to return a specific error of when a user is not registered on the system
var ErrUserNotFound = errors.New("jwtbackend: User not found")

//GetJWTAuthenticationBackend returns singleton instance
func GetJWTAuthenticationBackend() (*jwtAuthenticationBackend, error) {
	once.Do(func() {

		privateKey, err := getPrivateKey()
		if err != nil {
			log.Println(err)
			return
		}

		publicKey := privateKey.PublicKey

		authBackendInstance = &jwtAuthenticationBackend{privateKey, &publicKey}

	})

	if authBackendInstance != nil {
		return authBackendInstance, nil
	}

	return nil, errors.New("Unable to set private and public keys")
}

func (backend *jwtAuthenticationBackend) GenerateToken(userUUID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = userUUID
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func (backend *jwtAuthenticationBackend) Authenticate(user *model.LoginInfo) (uuid.UUID, error) {

	instance, err := database.GetUserDatabaseManagerInstance()
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}

	checkUser, err := instance.GetUserByEmail(user.Username)
	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return uuid.Nil, ErrUserNotFound
		}
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.HashCode), []byte(user.Password))
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}

	return checkUser.UUID, nil
}

func (backend *jwtAuthenticationBackend) getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

func (backend *jwtAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
	/* redisConn := redis.Connect()
	return redisConn.SetValue(tokenString, tokenString, backend.getTokenRemainingValidity(token.Claims["exp"])) */
	return nil
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(settings.Get().GetPrivateKeyPath());
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	size := pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return privateKeyImported, nil
}

/*func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().GetPublicKeyPath())
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}*/
