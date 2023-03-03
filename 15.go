package main

import "fmt"

func createHugeString(n int) string {
	return "какая-то строка"
}

// не очень поятна операция побитового смещения зачем тут,но один из вариантов ,что мы предполагаем взять допустим 4 элемента из строки v[:4],так как строка это массив байт,
//а символы могут быть в кодировке UTF-8 и могуть занимать больше чем один байт

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	justString = v[:4]

	fmt.Println(justString)
}

func main() {
	someFunc()
}
