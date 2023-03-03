package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1) посылаем горутине в канал значение,при котором горутина заврешает свою работу
func stoproutine1(ch chan int, wg *sync.WaitGroup) {
	for {
		print()
		select {
		case <-ch:
			wg.Done()
			return
		default:
			print("горутина пашет")
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	////1 вариант
	ch := make(chan int)
	wg.Add(1)
	go stoproutine1(ch, &wg)
	time.Sleep(5 * time.Second)
	ch <- 1
	wg.Wait()

	//2 вариант
	context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("контекст передал сообщение через канал и закрывает горутину")
				wg.Done()
				return
			default:
				fmt.Println("что-то делаю")
			}
		}
	}(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()

	//3 воспользуемся mutex и будем в основном канале увеличивать до 100 , а в другой горутине проверять условие <99,
	//если не выполняется,то завершаем работу горутины
	mu := sync.RWMutex{}
	i := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			mu.RLock()
			fmt.Println(i)
			if i > 99 {
				mu.RUnlock()
				break
			}
			mu.RUnlock()
		}
		fmt.Println("работаем")
	}()

	for j := 0; j < 100; j++ {
		mu.Lock()
		i++
		mu.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
}
