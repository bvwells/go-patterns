package behavioral

// Memento stores the state of the Number.
type Memento struct {
	state int
}

// NewMemento creates a new memento.
func NewMemento(value int) *Memento {
	return &Memento{value}
}

// Number represents an integer which can be operated on.
type Number struct {
	value int
}

// NewNumber creates a new Number.
func NewNumber(value int) *Number {
	return &Number{value}
}

// Dubble doubles the value of the number.
func (n *Number) Dubble() {
	n.value = 2 * n.value
}

// Half halves the value of the number.
func (n *Number) Half() {
	n.value = n.value / 2
}

// Value returns the value of the number.
func (n *Number) Value() int {
	return n.value
}

// CreateMemento creates a Memento with the current state of the number.
func (n *Number) CreateMemento() *Memento {
	return NewMemento(n.value)
}

// ReinstateMemento reinstates the value of the Number to the value of the memento.
func (n *Number) ReinstateMemento(memento *Memento) {
	n.value = memento.state
}
