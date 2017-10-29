package structural

import (
	"testing"

	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestNewBasicTimeImp_ReturnsNonNil(t *testing.T) {
	impl := NewBasicTimeImp(1, 2)
	assert.NotNil(t, impl)
}

func TestBasicTimeImplTell_TellsTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewBasicTimeImp(1, 2)
	impl.Tell()
	assert.Equal(t, "The time is 01:02\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewCivilianTimeImp_ReturnsNonNil(t *testing.T) {
	impl := NewCivilianTimeImp(1, 2, false)
	assert.NotNil(t, impl)
}

func TestCivilianTimeImpTell_TellsAMTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewCivilianTimeImp(1, 2, false)
	impl.Tell()
	assert.Equal(t, "The time is 01:02 AM\n", outputWriter.(*bytes.Buffer).String())
}

func TestCivilianTimeImpTell_TellsPMTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewCivilianTimeImp(1, 2, true)
	impl.Tell()
	assert.Equal(t, "The time is 01:02 PM\n", outputWriter.(*bytes.Buffer).String())
}

func TestNewZuluTimeImp_ReturnsNonNil(t *testing.T) {
	impl := NewZuluTimeImp(1, 2, 5)
	assert.NotNil(t, impl)
}

func TestZuluTimeImpTell_TellsEasternStandardTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewZuluTimeImp(1, 2, 5)
	impl.Tell()
	assert.Equal(t, "The time is 01:02 Eastern Standard Time\n", outputWriter.(*bytes.Buffer).String())
}

func TestZuluTimeImpTell_TellsCentralStandardTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewZuluTimeImp(1, 2, 6)
	impl.Tell()
	assert.Equal(t, "The time is 01:02 Central Standard Time\n", outputWriter.(*bytes.Buffer).String())
}

func TestZuluTimeImpTell_TellsTime(t *testing.T) {
	bufferOutputWriter := outputWriter
	outputWriter = new(bytes.Buffer)
	defer func() { outputWriter = bufferOutputWriter }()

	impl := NewZuluTimeImp(1, 2, 7)
	impl.Tell()
	assert.Equal(t, "The time is 01:02 \n", outputWriter.(*bytes.Buffer).String())
}

func TestTell_TellsTime(t *testing.T) {
	times := make([]*Time, 3, 3)
	times[0] = NewTime(14, 30)
	times[1] = NewCivilianTime(2, 30, true)
	times[2] = NewZuluTime(14, 30, 6)

	for i := 0; i < len(times); i++ {
		times[i].Tell()
	}
}
