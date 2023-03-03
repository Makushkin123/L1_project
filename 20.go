package main

import (
	"fmt"
	"strings"
)

// функция делает реверс строки, используем метод Split для получения слайса string и потом итерируемся и меняем элементы слайса
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun moon cat cat cat light day night"
	fmt.Println(reverseWords(s))
}
