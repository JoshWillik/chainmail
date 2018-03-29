package chainmail

type Feed interface{
	Open(chan Message) error
	Close()
}
