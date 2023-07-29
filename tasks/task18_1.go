// 18. Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const goroutines = 10_000

type AtomicCounter struct {
	Value int32
}

func main() {
	var counter AtomicCounter

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter:", counter.Value)
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt32(&c.Value, 1)
}
