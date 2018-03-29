package chainmail

import (
	"testing"
)

type TestFeed struct{
	messages chan Message
	Messages []Message
}

func (f TestFeed) Open(m chan Message) error {
	f.messages = m
	go func(){
		for _, message := range f.Messages {
			m <- message
		}
		f.Close()
	}()
	return nil
}

func (f TestFeed) Close() {
	close(f.messages)
}

type testHandler struct{
	callback func(Message) error
}
func (s testHandler) ProcessMessage(m Message) error {
	return s.callback(m)
}

func makeHandler(fn func(Message) error) testHandler {
	return testHandler{fn}
}

func TestPipe(t *testing.T){
	processed := make([]Message, 0)
	testFeed := TestFeed{
		Messages: []Message{
			{Subject: "test 1"},
		},
	}
	pipe := Pipe{
		Feed: testFeed,
		Handler: makeHandler(func (m Message) error {
			processed = append(processed, m)
			return nil
		}),
	}
	done := make(chan error)
	go func(){
		done <- pipe.ProcessMessages()
	}()
	if err := <-done; err != nil {
		t.Error(err)
		t.Fail()
	}
	if len(processed) != 1 {
		t.Errorf("%d messages were processed", len(processed))
		for i, msg := range processed {
			t.Errorf("- %d: %s", i, msg.Subject)
		}
		t.Fail()
	}
}
