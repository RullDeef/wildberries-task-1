// 14. Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа
// interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var mystery interface{} = make(chan map[string]func())

	switch v := mystery.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	default:
		// check chan type
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Chan {
			fmt.Printf("channel %T\n", v)
		} else {
			fmt.Println("unknown type")
		}
	}
}
