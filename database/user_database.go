package database

import (
	"database/sql"
	"log"

	"sync"

	"errors"

	"github.com/FEUPTalks/Backend/model"

	"github.com/FEUPTalks/Backend/model/roles/roleFactory"
	//loading the driver anonymously, aliasing its package qualifier so none of its exported names are visible to our code
	_ "github.com/go-sql-driver/mysql"
)

//userDatabaseManager used to manage the talk_store
type userDatabaseManager struct {
	database *sql.DB
}

const (
	uDriverName       string = "mysql"
	uConnectionString string = "lesteamb:99RedBalloons@tcp(127.0.0.1:3306)/user_store?parseTime=true"
)

var uInstance *userDatabaseManager
var uOnce sync.Once

//GetUserDatabaseManagerInstance returns singleton uInstance
func GetUserDatabaseManagerInstance() (*userDatabaseManager, error) {
	uOnce.Do(func() {
		var db *sql.DB
		var err error

		db, err = sql.Open(uDriverName, uConnectionString)
		if err != nil {
			db.Close()
			log.Fatal(err)
		}
		uInstance = &userDatabaseManager{db}
	})
	if uInstance != nil {
		return uInstance, nil
	}
	return nil, errors.New("Unable to create access to the database")
}

func (manager *userDatabaseManager) CloseConnection() (err error) {
	err = manager.database.Close()
	if err != nil {
		log.Println(err)
	}
	return
}

func (manager *userDatabaseManager) Ping() error {
	err := manager.database.Ping()
	if err != nil {
		log.Println(err)
		return errors.New("Unable to access database")
	}
	return nil
}

//GetUser retrieves users with specific id from the database
func (manager *userDatabaseManager) GetUserByEmail(userEmail string) (*model.User, error) {
	stmt, err := manager.database.Prepare("select * from user where Email = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var user = model.NewUser()
	var roleTemp uint8

	err = stmt.QueryRow(userEmail).Scan(
		&user.UUID,
		&user.Email,
		&user.Name,
		&user.HashCode,
		&roleTemp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tempRole, err := roleFactory.GetRole(roleTemp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user.SetRole(tempRole)

	return user, nil
}
