// 7. Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	dataMap := make(map[int]int)

	// использование мютекса для обеспечения
	// монопольного доступа к мапе при записи
	var mutex sync.Mutex

	var wg sync.WaitGroup
	for min := 0; min < 3000; min += 100 {
		wg.Add(1)
		max := min + 100
		go func(min, max int) {
			writerLoop(dataMap, &mutex, min, max)
			wg.Done()
		}(min, max)
	}

	wg.Wait()
	fmt.Println("map len:", len(dataMap))
}

func writerLoop(dataMap map[int]int, mutex *sync.Mutex, min, max int) {
	for i := min; i < max; i++ {
		mutex.Lock()
		dataMap[i] = i
		mutex.Unlock()
	}
}
