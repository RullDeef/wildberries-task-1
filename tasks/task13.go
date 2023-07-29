// 13. Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 14
	b := 27
	fmt.Println(a, b)

	// 1 способ
	a, b = b, a
	fmt.Println(a, b)

	// 2 способ
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)

	// 3 способ
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	// 4 способ
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}
