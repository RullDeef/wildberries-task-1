// 25. Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting to sleep for 3s")
	time.Sleep(3 * time.Second)
	fmt.Println("woke up")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
