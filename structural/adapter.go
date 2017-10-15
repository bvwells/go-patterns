package structural

import (
	"fmt"
)

// Target is the target interface to adapt to.
type Target interface {
	Execute()
}

// Adaptee defines the old interface which needs to be adapted.
type Adaptee struct {
}

// SpecificExecute defines the old interface for executing which
// needs to be adapted to the new interface defined by Target.
func (a *Adaptee) SpecificExecute() {
	fmt.Fprint(outputWriter, "Executing SpecificExecute.")
}

// Adapter is the adaptor to the new interface Target.
type Adapter struct {
	*Adaptee
}

// Execute adapts from the new interface Execute to the old SpecificExecute.
func (a *Adapter) Execute() {
	a.SpecificExecute()
}
