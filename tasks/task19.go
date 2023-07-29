// 19. Разработать программу, которая переворачивает подаваемую на
// вход строку (например: «главрыба — абырвалг»). Символы могут
// быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	str = reverse(str)

	fmt.Println(str)
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
