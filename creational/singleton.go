package creational

import (
    "sync"
)

// Singleton deacribes a struct of which only a single
// instance can exist.
type Singleton struct {
}

var instance *Singleton
var once sync.Once

// GetInstance ensures that only a single instance of the
// struct singleton gets created and provides a global point
// to access to it.
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
