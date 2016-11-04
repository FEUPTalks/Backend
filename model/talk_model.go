package model

import "time"

//Talk struct
type Talk struct {
	TalkID               string    `json:"talkID"`
	Title                string    `json:"title"`
	Summary              string    `json:"summary"`
	ProposedInitialDate  time.Time `json:"proposedInitialDate"`
	ProposedEndDate      time.Time `json:"proposedEndDate"`
	DefinitiveDate       time.Time `json:"definitiveDate"`
	Duration             uint8     `json:"duration"`
	ProponentName        string    `json:"proponentName"`
	ProponentEmail       string    `json:"proponentEmail"`
	ProponentAffiliation string    `json:"proponentAffiliation"`
	SpeakerName          string    `json:"speakerName"`
	SpeakerBrief         string    `json:"speakerBrief"`
	SpeakerAffiliation   string    `json:"speakerAffiliation"`
	HostName             string    `json:"hostName"`
	HostEmail            string    `json:"hostEmail"`
	Snack                string    `json:"snack"`
	Room                 string    `json:"room"`
}
