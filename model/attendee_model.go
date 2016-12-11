package model

type Attendee struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	TalkID           int    `json:"talkID"`
	IsAttendingSnack bool   `json:"isAttendingSnack"`
}

//Creates a new empty Talk Registration
func NewTalkRegistration() *Attendee {
	talkRegistration := &Attendee{}
	return talkRegistration
}
