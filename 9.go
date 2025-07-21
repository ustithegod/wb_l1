package main

import (
	"fmt"
	"time"
)

const LENGTH = 100

func generateArray() (arr [LENGTH]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return
}

func firstWorker(arr [LENGTH]int, firstCh chan int) {
	defer close(firstCh)

	for _, v := range arr {
		firstCh <- v
	}
}

func secondWorker(firstCh <-chan int, secondCh chan int) {
	defer close(secondCh)

	for v := range firstCh {
		secondCh <- v * 2
	}
}

func thirdWorker(secondCh <-chan int) {
	for v := range secondCh {
		fmt.Println(v)
	}
}

func main() {
	array := generateArray()

	firstCh := make(chan int, LENGTH)
	secondCh := make(chan int, LENGTH)

	go firstWorker(array, firstCh)
	go secondWorker(firstCh, secondCh)
	go thirdWorker(secondCh)

	time.Sleep(5 * time.Second)

}
