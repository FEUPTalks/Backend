package talkState

//WaitingTalkState state 6 of a talk
type WaitingTalkState struct {
}

//Handle implementation TalkState interface
func (*WaitingTalkState) Handle() uint8 {
	return 6
}
