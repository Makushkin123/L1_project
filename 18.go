package main

import (
	"fmt"
	"sync"
)

type Count struct {
	sync.Mutex
	count int
}

func add(c *Count, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	c.count++
	c.Unlock()
}

func main() {

	x := Count{count: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			add(&x, &wg)
		}()
	}

	wg.Wait()
	fmt.Println(x.count)

}
