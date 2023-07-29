// 16. Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

func main() {
	numbers := []int{86, 50, 10, 48, 37, 13, 64, 49, 71, 29, 43, 18, 52, 48, 17, 39}

	quickSort(numbers)
	fmt.Println(numbers)
}

func quickSort(numbers []int) {
	if edgeCase(numbers) {
		return
	}

	pivot := pivot(numbers)
	i, j := 0, len(numbers)-1
	for {
		for numbers[i] < pivot {
			i++
		}
		for numbers[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		numbers[i], numbers[j] = numbers[j], numbers[i]
		i++
		j--
	}

	quickSort(numbers[:i])
	quickSort(numbers[i:])
}

func edgeCase(numbers []int) bool {
	if len(numbers) == 2 && numbers[0] > numbers[1] {
		numbers[0], numbers[1] = numbers[1], numbers[0]
	}
	return len(numbers) <= 2
}

func pivot(numbers []int) int {
	min, max := numbers[0], numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return (min + max) / 2
}
