package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func channelGoroutine(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("горутина 1 работает")
	for {
		select {
		case <-stop:
			fmt.Println("\nгорутина 1 завершена с помощью канала")
			return
		default:
			imitateWork()
		}
	}
}

func contextGoroutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("горутина 2 работает")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nгорутина 2 завершена с помощью контекста")
			return
		default:
			imitateWork()
		}
	}
}

func runtimeGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("горутина 3 работает")
	for {
		imitateWork()
		if time.Now().Second()%2 == 0 {
			fmt.Println("\nгорутина 3 завершена с помощью runtime.Goexit()")
			runtime.Goexit()
		}
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	// Способ 1. Оставнока с помощью канала уведомления
	ch1 := make(chan struct{})
	go channelGoroutine(ch1, &wg)
	time.AfterFunc(2*time.Second, func() {
		ch1 <- struct{}{}
	})

	wg.Wait()

	wg.Add(1)
	// Способ 2. Остановка с помощью контекста
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go contextGoroutine(ctx, &wg)

	wg.Wait()

	wg.Add(1)
	// Способ 3. Остановка с помощью runtime.Goexit()
	go runtimeGoroutine(&wg)

	wg.Wait()
}

func imitateWork() {
	fmt.Print(".")
	time.Sleep(700 * time.Millisecond)
}
