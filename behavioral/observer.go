package behavioral

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing

// Event describes an event to observe and notify on.
type Event struct {
	id string
}

// EventObserver describes an interface for observing events.
type EventObserver interface {
	OnNotify(event Event)
}

// observer is an implementation of the EventObserver interface.
type observer struct {
	name string
}

// NewEventObserver returns a new instance of an EventObserver.
func NewEventObserver(name string) EventObserver {
	return &observer{name}
}

// OnNotify logs the event being notified on.
func (o *observer) OnNotify(event Event) {
	fmt.Fprintf(outputWriter, "observer '%s' received event '%s'\n", o.name, event.id)
}

// EventNotifier describes an interface for registering and de-registering observers to
// be notified when an event occurs.
type EventNotifier interface {
	Register(obs EventObserver)
	Deregister(obs EventObserver)
	Notify(event Event)
}

// eventNotifer is an implementation of the EventNotifier interface.
type eventNotifer struct {
	observers []EventObserver
}

// NewEventNotifier returns a new instance of an EventNotifier.
func NewEventNotifier() EventNotifier {
	return &eventNotifer{}
}

// Register registers a new observer for notifying on.
func (e *eventNotifer) Register(obs EventObserver) {
	e.observers = append(e.observers, obs)
}

// Deregister de-registers an observer for notifying on.
func (e *eventNotifer) Deregister(obs EventObserver) {
	for i := 0; i < len(e.observers); i++ {
		if obs == e.observers[i] {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
		}
	}
}

// Notify notifies all observers on an event.
func (e *eventNotifer) Notify(event Event) {
	for i := 0; i < len(e.observers); i++ {
		e.observers[i].OnNotify(event)
	}
}
