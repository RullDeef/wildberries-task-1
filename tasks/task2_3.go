// 2. Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// пул горутин
	goroutinesCount := 3
	numberChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(goroutinesCount)
	for i := 0; i < goroutinesCount; i++ {
		go func() {
			for number := range numberChan {
				square := number * number
				fmt.Println("square of", number, "is", square)
			}
			wg.Done()
		}()
	}

	for _, number := range numbers {
		numberChan <- number
	}
	close(numberChan)

	// ожидание завершения запущенных горутин
	wg.Wait()
}
