package behavioral

import (
	"fmt"
)

// Machine defines a machine which can be swwitched on and off.
type Machine struct {
	current State
}

// NewMachine creates a new machine.
func NewMachine() *Machine {
	fmt.Fprintf(outputWriter, "Machine is ready.\n")
	return &Machine{NewOFF()}
}

// setCurrent sets the current state of the machine.
func (m *Machine) setCurrent(s State) {
	m.current = s
}

// On pushes the on button.
func (m *Machine) On() {
	m.current.On(m)
}

// Off pushes the off button.
func (m *Machine) Off() {
	m.current.Off(m)
}

// State describes the internal state of the machine.
type State interface {
	On(m *Machine)
	Off(m *Machine)
}

// ON describes the on button state.
type ON struct {
}

// NewON creates a new ON state.
func NewON() State {
	return &ON{}
}

// On does nothing.
func (o *ON) On(m *Machine) {
	fmt.Fprintf(outputWriter, "   already ON\n")
}

// Off switches the state from on to off.
func (o *ON) Off(m *Machine) {
	fmt.Fprintf(outputWriter, "   going from ON to OFF\n")
	m.setCurrent(NewOFF())
}

// OFF describes the off button state.
type OFF struct {
}

// NewOFF creates a new OFF state.
func NewOFF() State {
	return &OFF{}
}

// On switches the state from off to on.
func (o *OFF) On(m *Machine) {
	fmt.Fprintf(outputWriter, "   going from OFF to ON\n")
	m.setCurrent(NewON())
}

// Off does nothing.
func (o *OFF) Off(m *Machine) {
	fmt.Fprintf(outputWriter, "   already OFF\n")
}
