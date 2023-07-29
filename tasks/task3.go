// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
// (2^2+4^2...) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1 способ - использование атомарных операций
func way1() {
	numbers := []int64{2, 4, 6, 8, 10}
	var squareSum int64

	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int64) {
			square := number * number
			atomic.AddInt64(&squareSum, square)
			wg.Done()
		}(number)
	}
	wg.Wait()

	fmt.Println("way 1 -", squareSum)
}

// 2 способ - использование канала
func way2() {
	numbers := []int{2, 4, 6, 8, 10}

	// делаем буферизированный канал, который вместит
	// результаты всех горутин без блокировки
	squareChan := make(chan int, len(numbers))

	var wg sync.WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		go func(number int) {
			square := number * number

			// блокировка на записи не происходит
			squareChan <- square

			wg.Done()
		}(number)
	}
	wg.Wait()
	close(squareChan)

	var squareSum int
	// вычитываем все данные из закрытого канала
	for square := range squareChan {
		squareSum += square
	}

	fmt.Println("way 2 -", squareSum)
}

func main() {
	way1()
	way2()
}
