// 11. Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	set1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}
	set2 := map[int]struct{}{
		3: {},
		4: {},
		5: {},
	}

	intersection := intersect(set1, set2)
	fmt.Println(intersection)
}

func intersect[T comparable](set1, set2 map[T]struct{}) map[T]struct{} {
	intersection := make(map[T]struct{})

	for k := range set1 {
		if _, ok := set2[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return intersection
}
