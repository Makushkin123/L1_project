package main

import (
	"fmt"
	"time"
)

// записываем данные в канал ch
func writeToChannel(ch chan int) {
	var b byte
	i := 0

	for b = 250; b <= 255; b++ {
		ch <- i
		i++
	}
}

// считываем даднные из канала ch
func readFromChannel(ch chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}

func main() {

	//создаем канал для передачи даннных
	ch := make(chan int)
	var n time.Duration = 5

	go writeToChannel(ch)
	go readFromChannel(ch)

	// создает канал и по истечению времени передаем в канал time, после этого программа завершается
	<-time.After(time.Second * n)
	fmt.Println("время таймера истекло")

}
