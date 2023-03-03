package main

import (
	"fmt"
	"sync"
)

// создаем структуру map с rwmutex. RWmutex - механизм получения блокировки
type Map struct {
	mu    sync.RWMutex
	mappa map[int]int
}

// метод структуру Map, безопасно записывает значение в map
func (c *Map) Set(key int, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.mappa[key] = value

}

// //метод структуру Map, безопасно читает данные из map
func (c *Map) Get(key int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.mappa[key]
	// ключ не найден
	if !found {
		return -1, false
	}

	return value, true
}

// функция ,которая записывает новые значения в map
func writeToMap(mappa *Map, x1 int, x2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := x1; number <= x2; number++ {
		mappa.Set(number, number)
	}
}

// читаем все значения из map
func readToMap(mappa *Map, x1 int, x2 int) {
	for number := x1; number <= x2; number++ {
		res, found := mappa.Get(number)
		if !found {
			fmt.Println("нет такого ключа")
		}
		fmt.Println(number, ":", res)
	}
}

func main() {
	//механизм ожидания группы горутин
	wg := sync.WaitGroup{}
	x1 := 0
	x2 := 100

	mappa := &Map{
		mappa: map[int]int{},
	}

	//добавляем 2 задачи на ожидания
	wg.Add(2)
	go writeToMap(mappa, x1, x2, &wg)
	go writeToMap(mappa, x1+100, x2+100, &wg)

	//ждем пока все горутины исполнятся и завершаем работу программы
	wg.Wait()

	readToMap(mappa, x1, x2+100)

}
