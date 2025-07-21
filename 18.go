package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	counter int32
}

func (sc *SafeCounter) Add() {
	atomic.AddInt32(&sc.counter, 1)
}

func main() {
	var counter SafeCounter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(counter *SafeCounter, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Add()
			}
		}(&counter, &wg)
	}

	wg.Wait()

	fmt.Println("counter value is:", counter.counter)
}
