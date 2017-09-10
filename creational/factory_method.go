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

type Stooge interface {
	slap_stick()
}

type larry struct {
}

func (s *larry) slap_stick() {
	fmt.Print("Larry: poke eyes\n")
}

type moe struct {
}

func (s *moe) slap_stick() {
	fmt.Printf("Moe: slap head\n")
}

type curly struct {
}

func (s *curly) slap_stick() {
	fmt.Printf("Curly: suffer abuse\n")
}

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
