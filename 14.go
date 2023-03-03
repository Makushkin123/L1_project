package main

import "fmt"

// интерфейсное значение внутри хранит информацию о конкретном типе и его значение
func typeAssertion(value interface{}) {

	switch value.(type) {
	case bool:
		fmt.Println("тип bool")
	case string:
		fmt.Println("тип string")
	case int:
		fmt.Println("тип int")
	default:
		fmt.Println("тип channel")
	}
}

func main() {

	var x interface{} = 5

	typeAssertion(x)

}
