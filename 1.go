package main

import (
	"fmt"
	"strconv"
)

// структура human c полями name и year.Методы greet и live
type Human struct {
	name string
	year int
}

func (h Human) greet() {
	fmt.Println("hello,my name " + h.name)

}
func (h Human) live() {
	fmt.Println("я родился " + strconv.Itoa(h.year) + " назад")

}

// Структура Action, в которую встроена структура Human
type Action struct {
	Human
}

func main() {

	x := Human{name: "Artem", year: 22}

	//Вызываем методы структуры Human
	x.greet()
	x.live()

	a := Action{Human{name: "Artem", year: 22}}

	//Вызываем методы встроенной структуру у структуры Action
	//1)Вариант
	a.Human.greet()
	a.Human.live()

	//Более удобный вариант
	//2)
	a.greet()
	a.live()

}
