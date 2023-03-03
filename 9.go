package main

import (
	"fmt"
	"sync"
)

//передаем массив чисел по первому каналу
func massivToChannel(massiv []int, wg *sync.WaitGroup, ch chan int) {
	wg.Add(1)
	for _, value := range massiv {
		ch <- value
	}
	wg.Done()
	close(ch)
}

//слущаем и передаем массив * 2 по второму каналу
func massivToChannel2(massiv []int, wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	wg.Add(1)
	for value := range ch1 {
		ch2 <- value * 2
	}
	wg.Done()
	close(ch2)
}

func main() {

	wg := sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)

	massiv := []int{1, 2, 3, 4, 5}

	go massivToChannel(massiv, &wg, ch1)
	go massivToChannel2(massiv, &wg, ch1, ch2)

	//слушаем второй канал и передаем числа в stdout
	for value := range ch2 {
		fmt.Println(value)
	}

}
