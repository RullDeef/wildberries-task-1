// 22. Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2)
	a.Exp(a, big.NewInt(100), nil)
	fmt.Println("a:", a)

	b := big.NewInt(3)
	b.Exp(b, big.NewInt(60), nil)
	fmt.Println("b:", b)

	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Println("sum:", sum)

	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Println("sub:", sub)

	product := big.NewInt(0)
	product.Mul(a, b)
	fmt.Println("product:", product)

	div := big.NewInt(0)
	mod := big.NewInt(0)
	div.DivMod(a, b, mod)
	fmt.Println("div:", div)
	fmt.Println("mod:", mod)
}
