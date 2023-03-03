package main

import (
	"fmt"
	"strings"
)

// функция для проверки уникальности числа ,использую мапу для проверку для провреки уникальности букв
func uniqueWord(str string) bool {

	unique := make(map[rune]bool)

	lowerStr := strings.ToLower(str)
	fmt.Println(lowerStr)

	for _, value := range lowerStr {

		if _, flag := unique[value]; flag {
			return false
		}

		unique[value] = true
	}
	return true
}

func main() {

	fmt.Println(uniqueWord("AAbc"))

}
