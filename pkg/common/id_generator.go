package common

import "sync"

var once sync.Once
var instance *IDGenerator

// IDGenerator is a thread-safe generator of unique IDs. It is safe for concurrent use.
// Note: In case of scaling to multiple instances of the service,
// this generator should be replaced with a more scalable solution.

type IDGenerator struct {
	mu *sync.Mutex
	id int
}

func NewIDGenerator() *IDGenerator {
	// This is a thread-safe way to create a singleton instance of IDGenerator.
	once.Do(func() {
		instance = &IDGenerator{mu: &sync.Mutex{}, id: 0}
	})

	return instance
}

func (g *IDGenerator) NextID() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.id++
	return g.id
}
