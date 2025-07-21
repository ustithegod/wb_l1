package main

import (
	"fmt"
)

func reverseArray(arr []string) {
	i, j := 0, len(arr)-1
	for k := 0; k <= j/2+1; k++ {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	arr := []string{"sun", "dog", "snow"}
	reverseArray(arr)
	fmt.Println(arr)
}
