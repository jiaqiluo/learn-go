package counter

import "sync"

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

func NewCounter() *Counter {
	return &Counter{}
}
