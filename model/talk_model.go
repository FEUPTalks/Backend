package model

import (
	"errors"
)

//Talk struct
type Talk struct {
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	ProponentName  string `json:"proponentName"`
	ProponentEmail string `json:"proponentEmail"`
	HostName       string `json:"hostName"`
	HostEmail      string `json:"hostEmail"`
}

//NewEmptyTalk create a new Talk with no information
func NewEmptyTalk() *Talk {
	return &Talk{}
}

//NewCompleteTalk create a new Talk with complete information
func NewCompleteTalk(title string, summary string, proponentName string, proponentEmail string, hostName string, hostEmail string) *Talk {
	return &Talk{title, summary, proponentName, proponentEmail, hostName, hostEmail}
}

//GetTitle returns the Talk title
func (talk *Talk) GetTitle() string {
	return talk.Title
}

//SetTitle sets the talk's title
func (talk *Talk) SetTitle(title string) {
	talk.Title = title
}

//GetSummary returns the summary
func (talk *Talk) GetSummary() (string, error) {
	if len(talk.Summary) > 0 {
		return talk.Summary, nil
	}
	return "", errors.New("Summary not defined")
}

//SetSummary sets the talk's summary
func (talk *Talk) SetSummary(summary string) {
	talk.Summary = summary
}

//GetProponentName returns the proponent's name
func (talk *Talk) GetProponentName() string {
	return talk.ProponentName
}

//SetProponentName sets the proponent's name
func (talk *Talk) SetProponentName(summary string) {
	talk.Summary = summary
}

//GetProponentEmail returns the proponent's email
func (talk *Talk) GetProponentEmail() string {
	return talk.ProponentEmail
}

//GetHostName returns the host's name
func (talk *Talk) GetHostName() string {
	return talk.HostName
}

//GetHostEmail returns the host's email
func (talk *Talk) GetHostEmail() string {
	return talk.HostEmail
}
