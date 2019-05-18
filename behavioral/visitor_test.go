package behavioral

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	t.Parallel()

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
