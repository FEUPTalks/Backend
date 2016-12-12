package model

type TalkRegistration struct {
	Name                        string `json:"name"`
	Email                       string `json:"email"`
	TalkID                      int    `json:"talkID"`
	IsAttendingSnack            bool   `json:"isAttendingSnack"`
	WantsToReceiveNotifications bool   `json:"wantsToReceiveNotifications"`
}

//Creates a new empty Talk Registration
func NewTalkRegistration() *Attendee {
	talkRegistration := &Attendee{}
	return talkRegistration
}
