package database

import (
	"database/sql"
	"log"

	"sync"

	"errors"

	"github.com/FEUPTalks/Backend/model"

	"github.com/FEUPTalks/Backend/model/talkState"
	"github.com/FEUPTalks/Backend/model/talkState/talkStateFactory"
	//loading the driver anonymously, aliasing its package qualifier to so none of its exported names are visible to our code
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//TalkDatabaseManager used to manage the talk_store
type talkDatabaseManager struct {
	database *sql.DB
}

const (
	driverName       string = "mysql"
	connectionString string = "lesteamb:99RedBalloons@tcp(127.0.0.1:3306)/talk_store?parseTime=true"
)

var instance *talkDatabaseManager
var once sync.Once

//GetTalkDatabaseManagerInstance returns singleton instance
func GetTalkDatabaseManagerInstance() (*talkDatabaseManager, error) {
	once.Do(func() {
		var db *sql.DB
		var err error

		db, err = sql.Open(driverName, connectionString)
		if err != nil {
			db.Close()
			log.Fatal(err)
		}
		instance = &talkDatabaseManager{db}
	})
	if instance != nil {
		return instance, nil
	}
	return nil, errors.New("Unable to create access to the database")
}

func (manager *talkDatabaseManager) CloseConnection() (err error) {
	err = manager.database.Close()
	if err != nil {
		log.Println(err)
	}
	return
}

func (manager *talkDatabaseManager) Ping() error {
	err := manager.database.Ping()
	if err != nil {
		log.Println(err)
		return errors.New("Unable to access database")
	}
	return nil
}

//GetAllTalks retrieves all talks from the database
func (manager *talkDatabaseManager) GetAllTalks() ([]*model.Talk, error) {
	talks := make([]*model.Talk, 0)
	rows, err := manager.database.Query("select * from talk")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var talk = model.NewTalk()
		var stateTemp uint8
		err := rows.Scan(&talk.TalkID, &talk.Title, &talk.Summary,
			&talk.Date, &talk.DateFlex, &talk.Duration, &talk.ProponentName,
			&talk.ProponentEmail, &talk.SpeakerName, &talk.SpeakerBrief, &talk.SpeakerAffiliation,
			&talk.SpeakerPicture, &talk.HostName,
			&talk.HostEmail, &talk.Snack, &talk.Room, &talk.Other, &stateTemp)
		if err != nil {
			log.Println(err)
			continue
		}
		tempState, err := talkStateFactory.GetTalkState(stateTemp)
		if err != nil {
			log.Println(err)
			continue
		}
		talk.SetState(tempState)
		talks = append(talks, talk)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return talks, nil
}

//GetTalksWithState retrieves all talks with the given state from the database
func (manager *talkDatabaseManager) GetTalksWithState(state talkState.TalkState) ([]*model.Talk, error) {
	talks := make([]*model.Talk, 0)
	rows, err := manager.database.Query("select * from talk where state = ?", state.Handle())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var talk = model.NewTalk()
		var stateTemp uint8
		err := rows.Scan(&talk.TalkID, &talk.Title, &talk.Summary,
			&talk.Date, &talk.DateFlex, &talk.Duration, &talk.ProponentName,
			&talk.ProponentEmail, &talk.SpeakerName, &talk.SpeakerBrief, &talk.SpeakerAffiliation,
			&talk.SpeakerPicture, &talk.HostName,
			&talk.HostEmail, &talk.Snack, &talk.Room, &talk.Other, &stateTemp)
		if err != nil {
			log.Println(err)
			continue
		}
		tempState, err := talkStateFactory.GetTalkState(stateTemp)
		if err != nil {
			log.Println(err)
			continue
		}
		talk.SetState(tempState)
		talks = append(talks, talk)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return talks, nil
}

//GetTalk retrieves talks with specific id from the database
func (manager *talkDatabaseManager) GetTalk(talkID int) (*model.Talk, error) {
	stmt, err := manager.database.Prepare("select * from talk where talkID = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var talk = model.NewTalk()
	var stateTemp uint8

	err = stmt.QueryRow(talkID).Scan(&talk.TalkID, &talk.Title, &talk.Summary,
		&talk.Date, &talk.DateFlex, &talk.Duration, &talk.ProponentName,
		&talk.ProponentEmail, &talk.SpeakerName, &talk.SpeakerBrief, &talk.SpeakerAffiliation,
		&talk.SpeakerPicture, &talk.HostName,
		&talk.HostEmail, &talk.Snack, &talk.Room, &talk.Other, &stateTemp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	tempState, err := talkStateFactory.GetTalkState(stateTemp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	talk.SetState(tempState)

	return talk, nil
}

func (manager *talkDatabaseManager) SaveTalk(talk *model.Talk) error {
	stmt, err := manager.database.Prepare(
		`insert into talk (
			Title,
			Summary,
			Date,
			DateFlex,
			Duration,
			ProponentName,
			ProponentEmail,
			SpeakerName,
			SpeakerBrief,
			SpeakerAffiliation,
			SpeakerPicture,
			HostName,
			HostEmail,
			Snack,
			Room,
			Other,
			State) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(talk.Title, talk.Summary, talk.Date, talk.DateFlex, talk.Duration,
		talk.ProponentName, talk.ProponentEmail, talk.SpeakerName,
		talk.SpeakerBrief, talk.SpeakerAffiliation, talk.SpeakerPicture,
		talk.HostName, talk.HostEmail, talk.Snack, talk.Room, talk.Other, talk.GetStateValue())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//Returns all of the attendees that are registered in a given talk with ID == talkID
func (manager *talkDatabaseManager) GetTalkRegistrationsWithTalkID(talkID int) ([]*model.TalkRegistration, error) {
	talkRegistrations := make([]*model.TalkRegistration, 0)
	stmt, err := manager.database.Prepare("select * from talkRegistration where talkID = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(talkID)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var talkRegistration = model.NewTalkRegistration()
		err = rows.Scan(&talkRegistration.Email, &talkRegistration.TalkID, &talkRegistration.Name,
			&talkRegistration.IsAttendingSnack, &talkRegistration.WantsToReceiveNotifications)
		if err != nil {
			log.Println(err)
			continue
		}

		talkRegistrations = append(talkRegistrations, talkRegistration)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return talkRegistrations, nil
}

func (manager *talkDatabaseManager) GetTemporaryTalkRegistrationsWithTalkID(talkID int) ([]*model.TalkRegistration, error) {
	talkRegistrations := make([]*model.TalkRegistration, 0)
	stmt, err := manager.database.Prepare("select * from temporaryTalkRegistration where talkID = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(talkID)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var talkRegistration = model.NewTalkRegistration()
		err = rows.Scan(&talkRegistration.Email, &talkRegistration.TalkID, &talkRegistration.Name,
			&talkRegistration.IsAttendingSnack, &talkRegistration.WantsToReceiveNotifications)
		if err != nil {
			log.Println(err)
			continue
		}

		talkRegistrations = append(talkRegistrations, talkRegistration)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return talkRegistrations, nil
}

func (manager *talkDatabaseManager) EditTalkRegistration(talkID int, email string, talkRegistration *model.TalkRegistration) error {
	stmt, err := manager.database.Prepare(
		`update talkRegistration set
			Name=?,
			IsAttendingSnack=?,
			WantsToReceiveNotifications=? where talkID=? and email=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(talkRegistration.Name, talkRegistration.IsAttendingSnack,
		talkRegistration.WantsToReceiveNotifications, talkRegistration.TalkID, talkRegistration.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	err = deleteTemporaryTalkRegistration(talkID, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func deleteTemporaryTalkRegistration(talkID int, email string) error {
	stmt, err := instance.database.Prepare(`delete from temporaryTalkRegistration where email=? and talkID=?`)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := stmt.Exec(email, talkID)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (manager *talkDatabaseManager) SaveTalkRegistration(talkRegistration *model.TalkRegistration) error {
	stmt, err := manager.database.Prepare(
		`insert into talkRegistration (
			Email,
			TalkID,
			Name,
			IsAttendingSnack,
			WantsToReceiveNotifications) values (?,?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(talkRegistration.Email, talkRegistration.TalkID,
		talkRegistration.Name, talkRegistration.IsAttendingSnack, talkRegistration.WantsToReceiveNotifications)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//Adds a talk registration log
func (manager *talkDatabaseManager) SaveTalkRegistrationLog(talkRegistrationLog *model.TalkRegistrationLog) error {
	stmt, err := manager.database.Prepare(
		`insert into talkRegistrationLog (
			Email,
			TalkID,
			Name,
			IsAttendingSnack,
			WantsToReceiveNotifications,
			TransactionType,
			TransactionDate) values (?,?,?,?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(talkRegistrationLog.Email, talkRegistrationLog.TalkID,
		talkRegistrationLog.Name, talkRegistrationLog.IsAttendingSnack, talkRegistrationLog.WantsToReceiveNotifications,
		talkRegistrationLog.TransactionType, talkRegistrationLog.TransactionDate)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//Returns all of the attendees that are registered in a given talk with ID == talkID
func (manager *talkDatabaseManager) GetTalkRegistrationLogsWithTalkID(talkID int) ([]*model.TalkRegistrationLog, error) {
	talkRegistrationLogs := make([]*model.TalkRegistrationLog, 0)
	stmt, err := manager.database.Prepare("select * from talkRegistrationLog where talkID = ?")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(talkID)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var talkRegistrationLog = model.NewTalkRegistrationLog()
		err = rows.Scan(&talkRegistrationLog.LogID, &talkRegistrationLog.Name, &talkRegistrationLog.Email,
			&talkRegistrationLog.TalkID, &talkRegistrationLog.IsAttendingSnack, &talkRegistrationLog.WantsToReceiveNotifications,
			&talkRegistrationLog.TransactionType, &talkRegistrationLog.TransactionDate)
		if err != nil {
			log.Println(err)
			continue
		}

		talkRegistrationLogs = append(talkRegistrationLogs, talkRegistrationLog)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return talkRegistrationLogs, nil
}

//CheckIfRegistrationExistsInTalk
func (manager *talkDatabaseManager) CheckIfRegistrationExistsInTalk(email string, talkID int) (bool, error) {

	stmt, err := manager.database.Prepare("select * from talkregistration where email=? and talkID=?")
	if err != nil {
		log.Println(err)
		return false, err
	}

	talkRegistration := model.TalkRegistration{}

	err = stmt.QueryRow(email, talkID).Scan(&talkRegistration.Email, &talkRegistration.TalkID, &talkRegistration.Name,
		&talkRegistration.IsAttendingSnack, &talkRegistration.WantsToReceiveNotifications)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}

func (manager *talkDatabaseManager) CreateTemporaryTalkRegistrationChange(temporaryTalkRegistrationChange *model.TalkRegistration) error {
	stmt, err := manager.database.Prepare(
		`insert into temporaryTalkRegistration (
			Email,
			TalkID,
			Name,
			IsAttendingSnack,
			WantsToReceiveNotifications) values (?,?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(temporaryTalkRegistrationChange.Email, temporaryTalkRegistrationChange.TalkID,
		temporaryTalkRegistrationChange.Name, temporaryTalkRegistrationChange.IsAttendingSnack, temporaryTalkRegistrationChange.WantsToReceiveNotifications)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//SetTalk
func (manager *talkDatabaseManager) SetTalk(talk *model.Talk) error {
	stmt, err := manager.database.Prepare(`
	UPDATE Talk SET
		Title=?,
		Summary=?,
		Date=?,
		DateFlex=?,
		Duration=?,
		SpeakerName=?,
		SpeakerBrief=?,
		SpeakerAffiliation=?,
		SpeakerPicture=?,
		HostName=?,
		HostEmail=?,
		Snack=?,
		Room=?,
		Other=?
	WHERE TalkID=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(talk.Title, talk.Summary, talk.Date, talk.DateFlex, talk.Duration, talk.SpeakerName,
		talk.SpeakerBrief, talk.SpeakerAffiliation, talk.SpeakerPicture,
		talk.HostName, talk.HostEmail, talk.Snack, talk.Room, talk.Other, talk.TalkID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//SetTalkState
func (manager *talkDatabaseManager) SetTalkState(talkID int, state int) error {
	stmt, err := manager.database.Prepare(`
	UPDATE Talk SET
		State=?
	WHERE TalkID=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(state, talkID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//SetTalkRoom
func (manager *talkDatabaseManager) SetTalkRoom(talkID int, room string) error {
	stmt, err := manager.database.Prepare(`
	UPDATE Talk SET
		Room=?
	WHERE TalkID=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(room, talkID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (manager *talkDatabaseManager) SetTalkOther(talkID int, other string) error {
	stmt, err := manager.database.Prepare(`
	UPDATE Talk SET
		Other=?
	WHERE TalkID=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(other, talkID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (manager *talkDatabaseManager) SavePicture(filepath string) (int64, error) {
	stmt, err := manager.database.Prepare("insert into picture (filepath) values (?)")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	result, err := stmt.Exec(filepath)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

//GetPicture
func (manager *talkDatabaseManager) GetPictureByTalkID(id string) (string, error) {
	stmt, err := manager.database.Prepare("select speakerpicture from picture where pictureID IN (select speakerPicture from talk where talkID=?)")
	if err != nil {
		log.Println(err)
		return "", err
	}

	var picture string

	err = stmt.QueryRow(id).Scan(&picture)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return picture, nil
}

//DeleteLastTalk delete talk created in tests
func (manager *talkDatabaseManager) DeleteLastTalk() error {
	stmt, err := manager.database.Prepare(`DELETE FROM talk ORDER BY TalkID DESC LIMIT 1`)

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

//GetLastTalkID delete user created in tests
func (manager *talkDatabaseManager) GetLastTalkID() (int, error) {
	rows, err := manager.database.Query(`SELECT MAX(TalkID) FROM talk`, 1)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	var id int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return id, err
}

// Expire talks which already happened
func (manager *talkDatabaseManager) ExpireTalks() {
	talks, err := instance.GetAllTalks()
	if err != nil {
		log.Println(err)
		return
	}

	expire_time := time.Now().Add(24 * time.Hour).Local()

	for _,element := range talks {
		if element.StateValue != talkStateFactory.GetArchivedTalkStateValue() &&
			element.StateValue == talkStateFactory.GetPublishedTalkStateValue() {
			if !inTimeSpan(expire_time, element.Date) {
				instance.SetTalkState(element.TalkID, 6);
				log.Println("Expiring talk ", element.TalkID);
			}
		}
	}
}

func inTimeSpan(end, check time.Time) bool {
	return check.After(end)
}
