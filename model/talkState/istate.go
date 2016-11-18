package talkState

//TalkState used to apply the state design pattern
type TalkState interface {
	Handle() uint8
}
