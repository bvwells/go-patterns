package behavioral

import (
	"strconv"
)

// Handler defined a handler for handling a given handleID.
type Handler interface {
	Handle(handleID int) string
}

type handler struct {
	name     string
	next     Handler
	handleID int
}

// NewHandler returns a new Handler.
func NewHandler(name string, next Handler, handleID int) Handler {
	return &handler{name, next, handleID}
}

// Handle handles a given handleID.
func (h *handler) Handle(handleID int) string {
	if h.handleID == handleID {
		return h.name + " handled " + strconv.Itoa(handleID)
	}
	return h.next.Handle(handleID)
}
