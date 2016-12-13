package model

import "time"

type TalkRegistrationLog struct {
	LogID                       string    `json:"logID"`
	Name                        string    `json:"name"`
	Email                       string    `json:"email"`
	TalkID                      int       `json:"talkID"`
	IsAttendingSnack            bool      `json:"isAttendingSnack"`
	WantsToReceiveNotifications bool      `json:"wantsToReceiveNotifications"`
	TransactionType             int       `json:"transactionType"`
	TransactionDate             time.Time `json:"transactionDate"`
}

//Creates a new empty Talk Registration
func NewTalkRegistrationLog() *TalkRegistrationLog {
	talkRegistrationLog := &TalkRegistrationLog{}
	return talkRegistrationLog
}

//Creates a new Talk Registration Log with data from the talkRegistration and transactionType
func NewTalkRegistrationLogWithTalkRegistration(talkRegistration *TalkRegistration, transactionType int) *TalkRegistrationLog {
	talkRegistrationLog := &TalkRegistrationLog{}

	talkRegistrationLog.Name = talkRegistration.Name
	talkRegistrationLog.Email = talkRegistration.Email
	talkRegistrationLog.TalkID = talkRegistration.TalkID
	talkRegistrationLog.IsAttendingSnack = talkRegistration.IsAttendingSnack
	talkRegistrationLog.WantsToReceiveNotifications = talkRegistration.WantsToReceiveNotifications
	talkRegistrationLog.TransactionType = 0
	talkRegistrationLog.TransactionDate = time.Now()

	return talkRegistrationLog
}
