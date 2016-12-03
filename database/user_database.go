package database

import (
	"log"

	"github.com/FEUPTalks/Backend/model"

	//loading the driver anonymously, aliasing its package qualifier to so none of its exported names are visible to our code
	"github.com/FEUPTalks/Backend/model/roles/roleFactory"
	_ "github.com/go-sql-driver/mysql"
)

//GetUser retrieves users with specific id from the database
func (manager *talkDatabaseManager) GetUserByID(userID int) (*model.User, error) {
	stmt, err := manager.database.Prepare("select * from user where userID = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var user = model.NewUser()
	var roleTemp uint8

	err = stmt.QueryRow(userID).Scan(
		&user.UserID,
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
func (manager *talkDatabaseManager) GetUserByhash(hashcode string) (*model.User, error) {
	stmt, err := manager.database.Prepare("select * from user where hashcode = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var user = model.NewUser()
	var roleTemp uint8

	err = stmt.QueryRow(hashcode).Scan(
		&user.UserID,
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

func (manager *talkDatabaseManager) SaveUser(user *model.User) error {
	stmt, err := manager.database.Prepare(`insert into user (
		Email, 
		Name, 
		HashCode,
		RoleValue)
		values (?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(user.Email, user.Name, user.HashCode, user.RoleValue)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//SetUser
func (manager *talkDatabaseManager) SetUser(user *model.User) error {
	stmt, err := manager.database.Prepare(`
	UPDATE User SET 
		Email=?,
		Name=?,
		HashCode=?,
		RoleValue=?
	WHERE UserID=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(user.Email, user.Name, user.HashCode, user.RoleValue, user.UserID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//DeleteLastUser delete user created in tests
func (manager *talkDatabaseManager) DeleteLastUser() error {
	stmt, err := manager.database.Prepare(`DELETE FROM user ORDER BY UserID DESC LIMIT 1`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
