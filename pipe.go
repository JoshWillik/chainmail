package chainmail

type Pipe struct{
	Feed Feed
	Handler Handler
}

func (p Pipe) ProcessMessages() error {
	messages := make(chan Message, 10)
	if err := p.Feed.Open(messages); err != nil {
		return err
	}
	for message := range messages {
		p.Handler.ProcessMessage(message)
	}
	return nil
}
