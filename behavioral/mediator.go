package behavioral

import "fmt"

// WildStallion describes an interface for a Wild Stallion band member.
type WildStallion interface {
	SetMediator(mediator Mediator)
}

// Bill describes Bill S. Preston, Esquire.
type Bill struct {
	mediator Mediator
}

// SetMediator sets the mediator.
func (b *Bill) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

// Respond responds.
func (b *Bill) Respond() {
	fmt.Fprintf(outputWriter, "Bill: What?\n")
	b.mediator.Communicate("Bill")
}

// Ted describes Ted "Theodore" Logan.
type Ted struct {
	mediator Mediator
}

// SetMediator sets the mediator.
func (t *Ted) SetMediator(mediator Mediator) {
	t.mediator = mediator
}

// Talk talks through mediator.
func (t *Ted) Talk() {
	fmt.Fprintf(outputWriter, "Ted: Bill?\n")
	t.mediator.Communicate("Ted")
}

// Respond responds.
func (t *Ted) Respond() {
	fmt.Fprintf(outputWriter, "Ted: Strange things are afoot at the Circle K.\n")
}

// Mediator describes the interface for communicating between Wild Stallion band members.
type Mediator interface {
	Communicate(who string)
}

// ConcreateMediator describes a mediator between Bill and Ted.
type ConcreateMediator struct {
	Bill
	Ted
}

// NewMediator creates a new ConcreateMediator.
func NewMediator() *ConcreateMediator {
	mediator := &ConcreateMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

// Communicate communicates between Bill and Ted.
func (m *ConcreateMediator) Communicate(who string) {
	if who == "Ted" {
		m.Bill.Respond()
	} else if who == "Bill" {
		m.Ted.Respond()
	}
}
