// 17. Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	numbers := []int{2, 3, 3, 5, 7, 8, 11, 12, 21, 21, 133}

	fmt.Println(binarySearch(numbers, 8))
	fmt.Println(binarySearch(numbers, 3))
	fmt.Println(binarySearch(numbers, 10))
	fmt.Println(binarySearch(numbers, 21))
}

func binarySearch(numbers []int, number int) (int, bool) {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := (left + right) / 2
		if numbers[mid] > number {
			right = mid
		} else if numbers[mid] < number {
			left = mid + 1
		} else {
			return mid, true
		}
	}
	return left, false
}
