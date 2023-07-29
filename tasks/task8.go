// 8. Дана переменная int64. Разработать программу которая
// устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	flags := int64(0xFFFF0000)

	for i := 0; i < 64; i += 4 {
		if rand.Intn(2) == 0 {
			fmt.Println("set bit", i)
			flags = setBit(flags, i)
		} else {
			fmt.Println("reset bit", i)
			flags = resetBit(flags, i)
		}
	}

	fmt.Printf("flags: %b\n", flags)
}

func setBit(num int64, index int) int64 {
	return num | (1 << index)
}

func resetBit(num int64, index int) int64 {
	return num &^ (1 << index)
}
