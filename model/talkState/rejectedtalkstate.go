package talkState

//RejectedTalkState state 5 of a talk
type RejectedTalkState struct {
}

//Handle implementation TalkState interface
func (*RejectedTalkState) Handle() uint8 {
	return 2
}
