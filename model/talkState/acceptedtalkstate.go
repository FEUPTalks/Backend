package talkState

//AcceptedTalkState state 3 of a talk
type AcceptedTalkState struct {
}

//Handle implementation TalkState interface
func (*AcceptedTalkState) Handle() uint8 {
	return 3
}
