package chainmail

type Parallel struct{
	Handlers []Handler
}
func (p Parallel) ProcessMessage(m Message) error {
	errors := make(chan error)
	counter := 0
	for _, handler := range p.Handlers {
		go func(){
			if err := handler.ProcessMessage(m); err != nil {
				errors <- err
			}
			counter = counter+1
			if counter == len(p.Handlers) {
				close(errors)
			}
		}()
	}
	for err := range errors {
		return err
	}
	return nil
}
