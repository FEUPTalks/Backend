package talkState

//PublishedTalkState state 4 of a talk
type PublishedTalkState struct {
}

//Handle implementation TalkState interface
func (*PublishedTalkState) Handle() uint8 {
	return 4
}
