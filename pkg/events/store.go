package events

type EventStore struct {
	events []Event
}

func NewEventStore() *EventStore {
	return &EventStore{
		events: make([]Event, 0),
	}
}

func (s *EventStore) Append(event Event) {
	s.events = append(s.events, event)
}

func (s *EventStore) All() []Event {
	return s.events
}
