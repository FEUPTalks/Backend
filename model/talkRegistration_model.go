package model

type TalkRegistration struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	TalkID           int    `json:"talkID"`
	IsAttendingSnack bool   `json:"isAttendingSnack"`
}

//Creates a new empty Talk Registration
func NewTalkRegistration() *TalkRegistration {
	talkRegistration := &TalkRegistration{}
	return talkRegistration
}
