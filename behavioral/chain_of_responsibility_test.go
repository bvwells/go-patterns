package behavioral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler_ReturnsNonNil(t *testing.T) {
	h := NewHandler("Barry", nil, 1)
	assert.NotNil(t, h)
}

func TestNewHandler_SetsName(t *testing.T) {
	name := "Barry"
	h := NewHandler(name, nil, 1)
	assert.NotNil(t, h.(*handler).name, name)
}

func TestNewHandler_SetsNext(t *testing.T) {
	next := handler{}
	h := NewHandler("Barry", &next, 1)
	assert.NotNil(t, h.(*handler).next, next)
}

func TestNewHandler_SetsRequestID(t *testing.T) {
	handleID := 1234
	h := NewHandler("Barry", nil, handleID)
	assert.NotNil(t, h.(*handler).handleID, handleID)
}

func TestHandle_FirstHandlerHandleRequest(t *testing.T) {
	barry := NewHandler("Barry", nil, 1)
	paul := NewHandler("Paul", barry, 2)
	expected := "Paul handled 2"
	actual := paul.Handle(2)
	assert.Equal(t, expected, actual)
}

func TestHandle_SecondHanlderHandlesRequest(t *testing.T) {
	barry := NewHandler("Barry", nil, 1)
	paul := NewHandler("Paul", barry, 2)
	expected := "Barry handled 1"
	actual := paul.Handle(1)
	assert.Equal(t, expected, actual)
}
