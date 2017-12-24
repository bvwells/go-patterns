package behavioral

import (
	"testing"
	//	"github.com/stretchr/testify/assert"
)

func TestVisitor(t *testing.T) {

	list := []Element{&This{}, &That{}, &TheOther{}}

	up := &UpVisitor{}
	down := &DownVisitor{}

	for i := 0; i < len(list); i++ {
		list[i].Accept(up)
	}
	for i := 0; i < len(list); i++ {
		list[i].Accept(down)
	}
}
