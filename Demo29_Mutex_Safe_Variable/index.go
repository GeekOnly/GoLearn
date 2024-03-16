package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter ios safe to use concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++ // Lock so only one goroutine at a time can access the map c.v.
	fmt.Printf(" c: %d \n", c.v[key])
	c.mux.Unlock()
}

// Value returns the returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock() // Lock so only goroutine at a time can access the map c.v.
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
