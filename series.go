package chainmail

type Series struct{
	Handlers []Handler
}
func (s Series) ProcessMessage(m Message) error {
	for _, handler := range s.Handlers {
		if err := handler.ProcessMessage(m); err != nil {
			return err
		}
	}
	return nil
}
