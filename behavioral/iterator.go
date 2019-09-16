package behavioral

// Iterator is an interface for an iterator.
type Iterator interface {

	// Index returns the index of the current iterator.
	Index() int

	// Value returns the current value of the iterator.
	Value() interface{}

	// HasNext returns whether another next element exists.
	HasNext() bool

	// Next increments the iterator to point to the next element.
	Next()
}

// ArrayIterator is an iterator which iterates over an array.
type ArrayIterator struct {
	array []interface{}
	index int
}

// Index returns the index of the current iterator.
func (i *ArrayIterator) Index() int {
	return i.index
}

// Value returns the current value of the iterator.
func (i *ArrayIterator) Value() interface{} {
	return i.array[i.index]
}

// HasNext returns whether another next element exists.
func (i *ArrayIterator) HasNext() bool {
	return i.index+1 != len(i.array)
}

// Next increments the iterator to point to the next element.
func (i *ArrayIterator) Next() {
	if i.HasNext() {
		i.index++
	}
}
