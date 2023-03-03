package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

// читаем из канала и выводим сообщение в stdout
func readChannel(ch chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {

	flag := true
	n := 10
	var b byte
	ch := make(chan int)

	//Буферезированный канал для записи сигнала ОС
	sigch := make(chan os.Signal, 1)

	// если мы нажмем CTL + C,то поступит сигнал в канал от операционной системы
	signal.Notify(sigch, os.Interrupt)

	for i := 0; i < n; i++ {
		go readChannel(ch)
	}

	//бесконечный цикл,который генерирует целые числа и записывает в канал ch
	for b = 250; b <= 255; b++ {
		if flag == false {
			break
		}
		select {
		case <-sigch:
			//graceful shutdown заключается в том, что мы закрываем безопасно наш канал
			close(ch)
			fmt.Println("close")
			flag = false
		default:
			ch <- rand.Int()
		}
	}

}
