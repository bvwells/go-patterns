package creational

import (
	"fmt"
)

type StoogeType int

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
	fmt.Print("Larry: Poke eyes\n")
}

type moe struct {
}

func (s *moe) SlapStick() {
	fmt.Printf("Moe: Slap head\n")
}

type curly struct {
}

func (s *curly) SlapStick() {
	fmt.Printf("Curly: Suffer abuse\n")
}

// NewStooge creates new stooges given the stooge type.
// Nil is returned if the stooge type is not recognised.
func NewStooge(stooge StoogeType) Stooge {
	if stooge == Larry {
		return &larry{}
	} else if stooge == Moe {
		return &moe{}
	} else if stooge == Curly {
		return &curly{}
	}
	return nil
}
