package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
)

func initWorkers(n int, ch chan int) {
	for i := 0; i < n; i++ {
		go worker(i, ch)
	}
}

func worker(id int, ch <-chan int) {
	for v := range ch {
		fmt.Printf("worker %d read: '%d'\n", id, v)
	}
}

func main() {
	var numberOfWorkers int
	fmt.Print("enter the number of workers: ")
	fmt.Scanln(&numberOfWorkers)
	if numberOfWorkers < 1 {
		fmt.Println("number of workers must be greater than 0")
		return
	}

	ch := make(chan int, numberOfWorkers)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	initWorkers(numberOfWorkers, ch)

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}
