package main

import "fmt"

// функция возращает квадрат числа в канал
func squareSum(value int, ch chan int) {
	res := value * value
	ch <- res
}

func main() {

	massiv := []int{2, 4, 6, 8, 10}
	//создаем канал для записи результатов функции squareSum из goroutine
	ch := make(chan int)
	sum := 0

	for _, value := range massiv {
		go squareSum(value, ch)
	}

	// читаем данные из канала и вычисляем общую сумму
	for range massiv {
		sum += <-ch
	}

	fmt.Println(sum)
}
