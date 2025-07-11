package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, value := range array {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			fmt.Println(number * number)
		}(value)
	}

	wg.Wait()
}
