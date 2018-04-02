package chainmail

type Feed interface{
  Init() error
	Open(chan Message) error
	Close()
}
