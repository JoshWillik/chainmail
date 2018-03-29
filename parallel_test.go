package chainmail

import (
	"time"
	"testing"
)

func TestParallelHandler(t *testing.T){
	processed := 0
	waiter := makeHandler(func (m Message) error {
		time.Sleep(time.Second/10)
		processed = processed + 1
		return nil
	})
	handler := Parallel {
		Handlers: []Handler{
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
			waiter,
		},
	}
	start := time.Now()
	if err := handler.ProcessMessage(Message{}); err != nil {
		t.Error(err)
		t.FailNow()
	}
	executionTime := time.Now().Sub(start)
	if executionTime > 2*time.Second/10 {
		t.Errorf("execution took %s, should be <0.2s",
			executionTime.Truncate(time.Second/10).String())
		t.FailNow()
	}
	if processed != 10 {
		t.Errorf("%d handlers were invoked out of 10 expected",
			processed)
		t.FailNow()
	}
}
