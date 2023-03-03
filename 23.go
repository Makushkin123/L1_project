package main

import "fmt"

// функция удаления i элемента в слайсе
func deleteValue(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {

	massiv := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := 5

	result := deleteValue(massiv, i)

	fmt.Println(result)
}
