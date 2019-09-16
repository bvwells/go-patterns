package creational

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout // modified during testing

// StoogeType is used as a enum for stooge types.
type StoogeType int

// Names of stooges as enums.
const (
	Larry StoogeType = iota
	Moe
	Curly
)

// Stooge provides an interface for interacting with stooges.
type Stooge interface {
	SlapStick()
}

type larry struct {
}

func (s *larry) SlapStick() {
	fmt.Fprint(outputWriter, "Larry: Poke eyes\n")
}

type moe struct {
}

func (s *moe) SlapStick() {
	fmt.Fprint(outputWriter, "Moe: Slap head\n")
}

type curly struct {
}

func (s *curly) SlapStick() {
	fmt.Fprint(outputWriter, "Curly: Suffer abuse\n")
}

// NewStooge creates new stooges given the stooge type.
// Nil is returned if the stooge type is not recognised.
func NewStooge(stooge StoogeType) Stooge {
	switch stooge {
	case Larry:
		return &larry{}
	case Moe:
		return &moe{}
	case Curly:
		return &curly{}
	}
	return nil
}
