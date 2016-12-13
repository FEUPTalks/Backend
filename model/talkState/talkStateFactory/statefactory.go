package talkStateFactory

import (
	"errors"

	"github.com/FEUPTalks/Backend/model/talkState"
)

//GetTalkState used to create TalkState implementations based on choice
func GetTalkState(state uint8) (talkState.TalkState, error) {
	switch state {
	case 1:
		return &talkState.ProposedTalkState{}, nil
	case 2:
		return &talkState.RejectedTalkState{}, nil
	case 3:
		return &talkState.AcceptedTalkState{}, nil
	case 4:
		return &talkState.PublishedTalkState{}, nil
	case 5:
		return &talkState.ArchivedTalkState{}, nil
	case 6:
		return &talkState.WaitingTalkState{}, nil
	default:
		return nil, errors.New("Requested state not available")
	}
}

//GetProposedTalkStateValue
func GetProposedTalkStateValue() uint8 {
	state := &talkState.ProposedTalkState{}
	return state.Handle()
}

//GetRejectedTalkStateValue
func GetRejectedTalkStateValue() uint8 {
	state := &talkState.RejectedTalkState{}
	return state.Handle()
}

//GetAcceptedTalkStateValue
func GetAcceptedTalkStateValue() uint8 {
	state := &talkState.AcceptedTalkState{}
	return state.Handle()
}

//GetPublishedTalkStateValue
func GetPublishedTalkStateValue() uint8 {
	state := &talkState.PublishedTalkState{}
	return state.Handle()
}

//GetArchivedTalkStateValue
func GetArchivedTalkStateValue() uint8 {
	state := &talkState.ArchivedTalkState{}
	return state.Handle()
}

//GetWaitingTalkStateValue
func GetWaitingTalkStateValue() uint8 {
	state := &talkState.WaitingTalkState{}
	return state.Handle()
}
