package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mu    *sync.RWMutex
}

func (c *counter) inc() {
	defer c.mu.Unlock()
	c.mu.Lock()

	c.count++
}

func (c *counter) value() int {
	defer c.mu.RUnlock()
	c.mu.RLock()

	return c.count
}

func main() {
	c := counter{
		mu: new(sync.RWMutex),
	}
	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(c.value())
}
