package main

import (
	"fmt"
	"math/big"
)

// сложение
func sum(x, y *big.Int) *big.Int {
	var z big.Int
	z.Add(x, y)
	return &z
}

// умножение
func multi(x, y *big.Int) *big.Int {
	var z big.Int
	z.Mul(x, y)
	return &z
}

// деление
func div(x, y *big.Int) *big.Int {
	var z big.Int
	z.Quo(x, y)
	return &z
}

// вычитание
func sub(x, y *big.Int) *big.Int {
	var z big.Int
	z.Neg(y)
	z.Add(x, &z)
	return &z
}

func main() {

	// воспользуемся пакетом big, для манипулирования числами > 2^20
	firstVariable := big.NewInt(500)
	secondVariable := big.NewInt(500)

	fmt.Println(firstVariable, ":", secondVariable)
	fmt.Println(sum(firstVariable, secondVariable))
	fmt.Println(sub(firstVariable, secondVariable))
	fmt.Println(multi(firstVariable, secondVariable))
	fmt.Println(div(firstVariable, secondVariable))

}
