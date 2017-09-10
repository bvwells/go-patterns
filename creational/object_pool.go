package creational

import "sync"

// PoolObject represents the object to be stored in the Pool.
type PoolObject struct {
}

// Pool represents the pool of objects to use.
type Pool struct {
	*sync.Mutex
	inuse     []*PoolObject
	available []*PoolObject
}

// NewPool creates a new pool.
func NewPool() *Pool {
	return &Pool{}
}

// Acquire acquires a new PoolObject to use from the pool.
// Here acquire creates a new instance of a PoolObject if none available.
func (p *Pool) Acquire() *PoolObject {
	p.Lock()
	var object *PoolObject
	if len(p.available) != 0 {
		object = p.available[0]
		p.available = append(p.available[:0], p.available[1:]...)
		p.inuse = append(p.inuse, object)
	} else {
		object := &PoolObject{}
		p.inuse = append(p.inuse, object)
	}
	p.Unlock()
	return object
}

// Release releases a PoolObject back to the Pool.
func (p *Pool) Release(object *PoolObject) {
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
