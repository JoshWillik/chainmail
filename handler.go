package chainmail

type Handler interface{
	ProcessMessage(Message) error
}
