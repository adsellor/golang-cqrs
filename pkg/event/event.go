package event

type Event struct {
	topic    string
	Callback func()
}

func NewEvent(topic string, callback func()) Event {
	return Event{
		topic:    topic,
		Callback: callback,
	}
}
