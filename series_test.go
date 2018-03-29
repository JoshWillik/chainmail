package chainmail

import (
	"testing"
)

func TestSeriesHandler(t *testing.T){
	processed := []string{}
	handler := Series{
		Handlers: []Handler{
			makeHandler(func (m Message) error {
				processed = append(processed, "from 1")
				return nil
			}),
			makeHandler(func (m Message) error {
				processed = append(processed, "from 2")
				return nil
			}),
		},
	}
	if err := handler.ProcessMessage(Message{}); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(processed) != 2 {
		t.Errorf("%d handlers were invoked out of 2 expected",
			len(processed))
		t.FailNow()
	}
	if processed[0] != "from 1" || processed[1] != "from 2" {
		t.Error("Unepected handler messages")
		for _, msg := range processed {
			t.Errorf("- %s", msg)
		}
		t.FailNow()
	}
}
