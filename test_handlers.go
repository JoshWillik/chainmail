package chainmail

type testHandler struct{
	callback func(Message) error
}
func (s testHandler) ProcessMessage(m Message) error {
	return s.callback(m)
}

func makeHandler(fn func(Message) error) testHandler {
	return testHandler{fn}
}
