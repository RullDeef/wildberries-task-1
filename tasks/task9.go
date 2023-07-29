// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2, после
// чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	firstChan := make(chan int)
	secondChan := make(chan int)

	go sender(firstChan, data)
	go multiplier(firstChan, secondChan, 2)
	printer(secondChan, "%d\n")
}

func sender(out chan<- int, data []int) {
	for _, i := range data {
		out <- i
	}
	close(out)
}

func multiplier(in <-chan int, out chan<- int, factor int) {
	for i := range in {
		out <- i * factor
	}
	close(out)
}

func printer(in <-chan int, format string) {
	for i := range in {
		fmt.Printf(format, i)
	}
}
