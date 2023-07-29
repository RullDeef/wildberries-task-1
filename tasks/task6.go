// 6. Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	// 1 способ - отправка сообщения о завершении в канал
	quitChan := make(chan struct{})
	go func(quitChan <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-quitChan:
				fmt.Println("gorutine quit (1)")
				return
			default:
				fmt.Println("gorutine work (1)")
				time.Sleep(1 * time.Second)
			}
		}
	}(quitChan)

	time.Sleep(2 * time.Second)
	quitChan <- struct{}{}

	// 2 способ - закрытие канала с данными
	dataChan := make(chan int)
	go func(dataChan <-chan int) {
		defer wg.Done()
		for data := range dataChan {
			fmt.Println("gorutine work (2): data =", data)
		}
		fmt.Println("gorutine quit (2)")
	}(dataChan)

	dataChan <- 3
	dataChan <- 2
	dataChan <- 1
	close(dataChan)

	// 3 способ - использование контекста
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("gorutine quit (3)")
				return
			default:
				fmt.Println("gorutine work (3)")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	cancel()
	wg.Wait()
}
