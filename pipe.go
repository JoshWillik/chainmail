package chainmail

type Pipe struct{
	Feed Feed
	Handler Handler
}

func (p Pipe) ProcessMessages() error {
	if err := p.Feed.Init(); err != nil {
		return err
	}
	messages := make(chan Message, 10)
	p.Feed.Open(messages)
	for message := range messages {
		p.Handler.ProcessMessage(message)
	}
	return nil
}
