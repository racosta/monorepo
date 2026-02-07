// Package sync provides a simple counter implementation using mutexes to ensure thread safety.
package sync

import "sync"

// Counter is a thread-safe counter implementation.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates and returns a new instance of Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter by 1 in a thread-safe manner.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter in a thread-safe manner.
func (c *Counter) Value() int {
	return c.value
}
