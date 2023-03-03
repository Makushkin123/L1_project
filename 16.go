package main

// алгоритм быстрой сортировки
func qwik_sort(massiv []int, low int, high int) []int {

	left := low
	right := high

	mid := massiv[(low+high)/2]

	for left <= right {
		//Ищем значения >
		for massiv[right] > mid {
			right--
		}
		//Ищем значения <
		for massiv[left] < mid {
			left++
		}
		if left <= right {
			massiv[right], massiv[left] = massiv[left], massiv[right]
			left++
			right--
		}
	}
	if right > low {
		qwik_sort(massiv, low, right)
	}
	if left > high {
		qwik_sort(massiv, left, high)
	}

	return massiv
}

func main() {
	lst := []int{1, 99, 100, 76, 23}
	low := lst[0]
	high := len(lst) - 1
	lst = qwik_sort(lst, low, high)
}
