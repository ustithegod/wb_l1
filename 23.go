package main

import "fmt"

// Решил попробовать дженерики, т.к. до этого никогда их не использовал

func deleteElement[T any](arr *[]T, idx int) {
	s := *arr
	copy(s[idx:], s[idx+1:])
	s = s[:len(s)-1]
	*arr = s
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	deleteElement(&arr, 2)
	fmt.Println(arr)
}
