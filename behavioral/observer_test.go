package behavioral

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventObserver_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	observer := NewEventObserver("peeping tom")
	assert.NotNil(t, observer)
}

func TestOnNotify_LogsMessage(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	event := Event{"something happened"}
	observer := NewEventObserver("peeping tom")
	observer.OnNotify(event)

	assert.Equal(t, "observer 'peeping tom' received event 'something happened'\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewEventNotifer_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	notifier := NewEventNotifier()
	assert.NotNil(t, notifier)
}

func TestRegister_RegistersObservers(t *testing.T) {
	t.Parallel()
	notifier := NewEventNotifier()
	observers := []EventObserver{
		NewEventObserver("tom"),
		NewEventObserver("dick"),
		NewEventObserver("harry"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}

	assert.Len(t, notifier.(*eventNotifer).observers, len(observers))
	for i := 0; i < len(observers); i++ {
		assert.Equal(t, observers[i], notifier.(*eventNotifer).observers[i])
	}
}

func TestDeregister_DeregistersObservers(t *testing.T) {
	t.Parallel()
	notifier := NewEventNotifier()
	observers := []EventObserver{
		NewEventObserver("tom"),
		NewEventObserver("dick"),
		NewEventObserver("harry"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}

	notifier.Deregister(observers[0])
	notifier.Deregister(observers[2])

	assert.Len(t, notifier.(*eventNotifer).observers, 1)
	assert.Equal(t, observers[1], notifier.(*eventNotifer).observers[0])
}

func TestNotify_NotifiesObservers(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	notifier := NewEventNotifier()
	observers := []EventObserver{
		NewEventObserver("tom"),
		NewEventObserver("dick"),
		NewEventObserver("harry"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}

	notifier.Notify(Event{"birthday!"})

	assert.Equal(t, "observer 'tom' received event 'birthday!'\n"+
		"observer 'dick' received event 'birthday!'\n"+
		"observer 'harry' received event 'birthday!'\n", outputWriter.(*bytes.Buffer).String())
}
