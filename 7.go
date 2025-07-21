package main

import (
	"fmt"
	"sync"
)

// Решил использовать Mutex, так как прочитал, что sync.Map оптимизирована для тех случаев,
// когда запись производится редко, а операций чтения много больше
type MutexMap struct {
	data map[string]int
	mu   sync.Mutex
}

func (mm *MutexMap) Set(key string, value int) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	mm.data[key] = value
}

func main() {
	m := &MutexMap{data: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(j int, wg *sync.WaitGroup) {
				defer wg.Done()
				m.Set(fmt.Sprint("key_%d", j), j)
			}(j, &wg)
		}
	}

	wg.Wait()
}
