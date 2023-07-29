// 26. Разработать программу, которая проверяет, что все символы
// в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// - abcd — true
// - abCdefAaf — false
// - aabcd — false

package main

import (
	"fmt"
	"unicode"
)

func main() {
	for _, str := range []string{"abcd", "abCdefAaf", "aabcd"} {
		fmt.Println(str, "-", unique(str))
	}
}

func unique(str string) bool {
	chars := make(map[rune]struct{})
	for _, ch := range str {
		ch = unicode.ToLower(ch)
		if _, ok := chars[ch]; ok {
			return false
		}
		chars[ch] = struct{}{}
	}
	return true
}
