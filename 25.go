package main

import "time"

// функция sleep,по истечению времени , time отправит сообщение в канал и главная gorytine ждет это сообщение
func sleep(second time.Duration) {
	<-time.After(second * time.Second)
}

func main() {
	var timer int64 = 10
	sleep(time.Duration(timer))
}
