package model

import (
	"time"

	"github.com/FEUPTalks/Backend/model/talkState"
	"github.com/FEUPTalks/Backend/model/talkState/talkStateFactory"
)

//Talk struct
type Talk struct {
	TalkID             int       `json:"talkID"`
	Title              string    `json:"title"`
	Summary            string    `json:"summary"`
	Date               time.Time `json:"date"`
	DateFlex           int       `json:"dateflex"`
	Duration           uint8     `json:"duration"`
	ProponentName      string    `json:"proponentName"`
	ProponentEmail     string    `json:"proponentEmail"`
	SpeakerName        string    `json:"speakerName"`
	SpeakerBrief       string    `json:"speakerBrief"`
	SpeakerAffiliation string    `json:"speakerAffiliation"`
	SpeakerPicture     int       `json:"speakerPicture"`
	HostName           string    `json:"hostName"`
	HostEmail          string    `json:"hostEmail"`
	Snack              int       `json:"snack"`
	Room               string    `json:"room"`
	Other              string    `json:"other"`
	StateValue         uint8     `json:"state"`
	state              talkState.TalkState
}

//NewTalk creates a new empty Talk
func NewTalk() *Talk {
	talkState, _ := talkStateFactory.GetTalkState(1)
	talk := &Talk{}
	talk.SetState(talkState)
	return talk
}

//GetStateValue returns value of the state the talk is in
func (talk *Talk) GetStateValue() uint8 {
	return talk.state.Handle()
}

//SetState changs the state of the talk
func (talk *Talk) SetState(state talkState.TalkState) {
	talk.state = state
	talk.StateValue = talk.state.Handle()
}
