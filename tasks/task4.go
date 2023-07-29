// 4. Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные
// из канала и выводят в stdout. Необходима возможность выбора количества
// воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	workersCount := inputWorkersCount()
	dataChan := make(chan int)

	// завершение воркеров с помощью контекста.
	// Воркеры могут выполнять долгие операции, которые необходимо
	// будет прервать по сигналу. Контекст позволит корректно
	// завершить всю цепочку вызовов.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// обработка сигнала прерывания
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		cancel()
	}()

	// запуск воркеров
	var wg sync.WaitGroup
	wg.Add(workersCount)
	for i := 1; i <= workersCount; i++ {
		go func(i int) {
			defer wg.Done()
			workerLoop(ctx, i, dataChan)
		}(i)
	}

	// главный поток генерирует сообщения
	for data := 0; ; data++ {
		dataChan <- data
		time.Sleep(200 * time.Millisecond)
		// проверка на прерывание
		if ctx.Err() != nil {
			close(dataChan)
			break
		}
	}

	wg.Wait()
	fmt.Println("main exited")
}

func workerLoop(ctx context.Context, id int, dataChan <-chan int) {
	for {
		select {
		case data := <-dataChan:
			fmt.Println(id, data)
		case <-ctx.Done():
			fmt.Printf("worker #%d exited\n", id)
			return
		}
	}
}

func inputWorkersCount() int {
	var workersCount int
	fmt.Print("input workers count: ")
	fmt.Scanln(&workersCount)

	if workersCount < 1 {
		fmt.Println("workers count must be > 0")
		os.Exit(1)
	}

	return workersCount
}
