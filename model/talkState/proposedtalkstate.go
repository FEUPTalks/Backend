package talkState

//ProposedTalkState state 1 of a talk
type ProposedTalkState struct {
}

//Handle implementation TalkState interface
func (*ProposedTalkState) Handle() uint8 {
	return 1
}
