package main

import (
	"fmt"
	"sync"
)

// функция для вывода в stdout квадрат числа
func square(value int, wg *sync.WaitGroup) {

	//defer используется для того, чтобы выполнить wg.Done() последним при выходе из функции при ее завершении
	//wg.Done уменьшает счетчик WaitGroup на 1
	defer wg.Done()

	res := value * value

	fmt.Println(res)
}

func main() {

	//WaitGroup ожидает завершения коллекции goroutine.
	wg := sync.WaitGroup{}

	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, value := range array {

		// метод add счетчик активных goroutine,которые мы будем ждать
		wg.Add(1)

		go square(value, &wg)
	}

	//блокирует исполнение потока до тех пор пока счетчик WaitGroup не обнулится.
	wg.Wait()

}
