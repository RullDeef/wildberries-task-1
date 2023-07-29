// 23. Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	slice := []string{"hello", "world", "and", "micro", "cat"}

	// 1 способ
	i := 2
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println(slice)

	// 2 способ
	for j := i; j < len(slice)-1; j++ {
		slice[j] = slice[j+1]
	}
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
