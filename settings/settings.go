package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const adminEmail string = "admin@teste.com"
const employee1Email string = "employee1@teste.com"
const employee2Email string = "employee2@teste.com"

var emails = [3]string{adminEmail, employee1Email, employee2Email}

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "../../settings/tests.json",
}

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

func (s Settings) GetPrivateKeyPath() string {
	absPath, _ := filepath.Abs(s.PrivateKeyPath)
	return absPath
}

func (s Settings) GetPublicKeyPath() string {
	absPath, _ := filepath.Abs(s.PublicKeyPath)
	return absPath
}

var settings Settings = Settings{}
var env = "preproduction"

//IsAdmin
func IsAdmin(email string) bool {
	return email == adminEmail
}

//IsAdminOrEmployee
func IsAdminOrEmployee(email string) bool {
	for _, temp := range emails {
		if temp == email {
			return true
		}
	}
	return false
}

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
