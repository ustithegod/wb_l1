package main

import (
	"fmt"
	"time"
)

func producer(ch chan struct{}) {
	for {
		ch <- struct{}{}
		fmt.Println("PRODUCER: sent message")
	}
}

func consumer(ch <-chan struct{}) {
	for range ch {
		fmt.Println("CONSUMER: read message")
	}
}

func main() {
	fmt.Print("enter the timeout (in seconds): ")
	var numberOfSeconds int
	fmt.Scanln(&numberOfSeconds)

	ch := make(chan struct{})
	go producer(ch)
	go consumer(ch)

	timeout := time.After(time.Duration(numberOfSeconds) * time.Second)

	<-timeout
	close(ch)
	fmt.Println("время работы истекло")
}
