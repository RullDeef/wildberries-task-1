// 15. К каким негативным последствиям может привести данный фрагмент
// кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//     v := createHugeString(1 << 10)
//     justString = v[:100]
// }
// func main() {
//     someFunc()
// }

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// При инициализации justString копирования не произойдет.
// Из-за этого вместо ожидаемых ~100 байт памяти на самом деле
// будет занято ~1 КБ.

// Чтобы этого не допустить, нужно явно скопировать необходимую
// часть строки v и позволить сборщику мусора очистить неиспользуемую
// память.

var justString string

func someFunc_OLD() string {
	v := createHugeString(1 << 10)
	justString = v[:100]
	return v // for testing only
}

func someFunc() string {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
	return v // for testing only
}

func main() {
	v1 := someFunc_OLD()
	fmt.Println("v1 share data with justString:",
		haveSameBuffer(v1, justString))

	v2 := someFunc()
	fmt.Println("v2 share data with justString:",
		haveSameBuffer(v2, justString))

	// вывод:
	// v1 share data with justString: true
	// v2 share data with justString: false
}

func createHugeString(size int) string {
	const pattern = "abcdefghigklmnopqrstuvwxyz"
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = pattern[i%len(pattern)]
	}
	return string(bytes)
}

func haveSameBuffer(s1, s2 string) bool {
	h1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	h2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	return h1.Data == h2.Data
}
