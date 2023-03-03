package main

import (
	"fmt"
	"sort"
)

// бинарный поиск для поиска элемента в массиве
func binarySearch(massiv []int, element int) bool {

	left := 0
	right := len(massiv) - 1

	for left <= right {
		middle := (left + right) / 2

		if massiv[middle] < element {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if left == len(massiv) || massiv[left] != element {
		return false
	}

	return true
}

func main() {
	massiv := []int{2, 4, 5, 1, 6, 7, 10, 9, 8}

	sort.Ints(massiv)

	res := binarySearch(massiv, -1)

	fmt.Println(res)
}
