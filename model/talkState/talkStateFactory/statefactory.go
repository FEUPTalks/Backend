package talkStateFactory

import (
	"errors"

	"github.com/RAyres23/LESTeamB-backend/model/talkState"
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
	default:
		return nil, errors.New("Requested state not available")
	}
}
