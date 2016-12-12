package talkState

//ArchivedTalkState state 3 of a talk
type ArchivedTalkState struct {
}

//Handle implementation TalkState interface
func (*ArchivedTalkState) Handle() uint8 {
	return 5
}
