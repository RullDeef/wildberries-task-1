// 5. Разработать программу, которая будет последовательно отправлять
// значения в канал, а с другой стороны канала — читать. По истечению
// N секунд программа должна завершаться.

package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 3 * time.Second

func main() {
	dataChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	// запускаем читателя в отдельной горутине.
	// Читатель завершается по закрытию канала с данными
	go func() {
		defer wg.Done()
		readerLoop(dataChan)
	}()

	timer := time.After(N)
writerLoop:
	for data := 0; ; data++ {
		select {
		case <-timer:
			close(dataChan)
			break writerLoop
		default:
			fmt.Println("write", data)
			dataChan <- data
			time.Sleep(200 * time.Millisecond)
		}
	}

	wg.Wait()
	fmt.Println("main(writer) exited")
}

func readerLoop(dataChan <-chan int) {
	for data := range dataChan {
		fmt.Println("read", data)
	}
	fmt.Println("reader exited")
}
