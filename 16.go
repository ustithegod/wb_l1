package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sort(arr []int, l, h int) int {
	pivot := arr[l]
	k := h
	for i := h; i > l; i-- {
		if arr[i] > pivot {
			swap(arr, k, i)
			k--
		}
	}

	swap(arr, k, l)

	return k
}

func quickSort(arr []int, l, h int) {
	if l < h {
		p := sort(arr, l, h)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, h)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
