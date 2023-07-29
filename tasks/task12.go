// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

package main

import "fmt"

func main() {
	strings := []string{
		"cat", "cat", "dog", "cat", "tree",
	}

	strMap := make(map[string]struct{})
	for _, str := range strings {
		strMap[str] = struct{}{}
	}

	fmt.Println(strMap)
}
