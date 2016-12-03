package model

type TalkRegistration struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	TalkID           int    `json:"talkID"`
	IsAttendingSnack bool   `json:"isAttendingSnack"`
}
