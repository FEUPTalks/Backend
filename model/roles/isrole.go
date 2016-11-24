package roles

//Role used to apply the state design pattern
type Role interface {
	Handle() uint8
}
