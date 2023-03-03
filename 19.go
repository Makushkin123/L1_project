package main

import "fmt"

// функция,которая переворачивает нашу строку
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {

	s := "日本語"

	res := reverse(s)
	fmt.Println(res)
}
