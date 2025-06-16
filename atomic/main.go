package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Atomicstructure struct {
	value int32
}

func (c *Atomicstructure) Increment() {
	atomic.AddInt32(&c.value, 1)
}

func (c *Atomicstructure) Display() int32 {
	return atomic.LoadInt32(&c.value)
}

func main() {
	count := Atomicstructure{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			count.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("The value of count.value is %v\n", count.Display())

}
