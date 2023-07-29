// 20. Разработать программу, которая переворачивает слова в строке.

// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	str = strings.Join(words, " ")

	fmt.Println(str)
}
