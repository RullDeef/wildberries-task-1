// 18. Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

const goroutines = 10_000

type AtomicCounter struct {
	value int
	mutex sync.Mutex
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

	fmt.Println("counter:", counter.value)
}

func (c *AtomicCounter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}
