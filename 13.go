package main

import "fmt"

func main() {

	x := 5
	y := 12
	fmt.Println(x, ",", y)

	//меняем числа местами,используя сумму и вычитание из суммы.
	x = x + y
	y = x - y
	x = x - y

	fmt.Println(x, ",", y)

}
