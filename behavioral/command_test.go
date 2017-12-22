package behavioral

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTalk(t *testing.T) {

	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	// Fred will "Execute" Barney which will result in a call to PassOn()
	// Barney will "Execute" Betty which will result in a call to Gossip()
	// Betty will "Execute" Wilma which will result in a call to Listen()
	wilma := NewPerson("Wilma", NewCommand(nil, nil))
	betty := NewPerson("Betty", NewCommand(&wilma, wilma.Listen))
	barney := NewPerson("Barney", NewCommand(&betty, betty.Gossip))
	fred := NewPerson("Fred", NewCommand(&barney, barney.PassOn))
	fred.Talk()

	assert.Equal(t, "Fred is talking.\n"+
		"Barney is passing on.\n"+
		"Betty is gossiping.\n"+
		"Wilma is listening.\n", outputWriter.(*bytes.Buffer).String())
}
