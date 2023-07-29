// 2. Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for _, number := range numbers {
		// каждая горутина вычисляет квадрат и выводит его на экран
		go func(number int) {
			square := number * number
			fmt.Println("square of", number, "is", square)
			wg.Done()
		}(number)
	}

	// ожидание завершения запущенных горутин
	wg.Wait()
}
