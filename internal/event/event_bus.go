package event

import (
	"sync"
)

type EventBus struct {
	subscribers map[string][]*EventHandler
	mutex       sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]*EventHandler),
	}
}

func (eb *EventBus) Publish(event Event) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	for _, subscriber := range eb.subscribers[event.topic] {
		go func(s *EventHandler) {
			s.Execute()
			s.channel <- event
		}(subscriber)
	}
}

func (eb *EventBus) Subscribe(event Event, subscriber *EventHandler) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	eb.subscribers[event.topic] = append(eb.subscribers[event.topic], subscriber)
}

func (eb *EventBus) Unsubscribe(event Event, subscriber *EventHandler) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	for i, s := range eb.subscribers[event.topic] {
		if s == subscriber {
			eb.subscribers[event.topic] = append(eb.subscribers[event.topic][:i], eb.subscribers[event.topic][i+1:]...)
			break
		}
	}
}
