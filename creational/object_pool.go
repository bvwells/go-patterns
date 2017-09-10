package creational

import (
	"sync"
)

// Pool represents the pool of objects to use.
type Pool struct {
	sync.Mutex
	inuse     []interface{}
	available []interface{}
	new       func() interface{}
}

// NewPool creates a new pool.
func NewPool(new func() interface{}) *Pool {
	return &Pool{new: new}
}

// Acquire acquires a new PoolObject to use from the pool.
// Here acquire creates a new instance of a PoolObject if none available.
func (p *Pool) Acquire() interface{} {
	p.Lock()
	var object interface{}
	if len(p.available) != 0 {
		object = p.available[0]
		p.available = append(p.available[:0], p.available[1:]...)
		p.inuse = append(p.inuse, object)
	} else {
		object = p.new()
		p.inuse = append(p.inuse, object)
	}
	p.Unlock()
	return object
}

// Release releases a PoolObject back to the Pool.
func (p *Pool) Release(object interface{}) {
	p.Lock()
	p.available = append(p.available, object)
	for i, v := range p.inuse {
		if v == object {
			p.inuse = append(p.inuse[:i], p.inuse[i+1:]...)
			break
		}
	}
	p.Unlock()
}
