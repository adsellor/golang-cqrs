package event

type EventHandler struct {
	channel chan Event
	event   Event
}

func NewEventHandler(event Event, eventBus *EventBus) *EventHandler {
	eh := &EventHandler{
		channel: make(chan Event),
		event:   event,
	}

	eventBus.Subscribe(event, eh)

	return eh
}

func (eh *EventHandler) Channel() chan Event {
	return eh.channel
}

func (eh *EventHandler) Topic() string {
	return eh.event.topic
}

func (eh *EventHandler) Execute() {
	eh.event.Callback()
}
