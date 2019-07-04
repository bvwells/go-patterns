package behavioral

import (
	"testing"
)

func TestIterator(t *testing.T) {
	t.Parallel()

	array := []interface{}{10.0, 20.0, 30.0, 40.0, 50.0}

	iterator := ArrayIterator{array, 0}

	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(float64)
		if value != array[index] {
			t.Errorf("Expected array value to equal %v, but received %v", array[index], value)
		}
	}
}
