package main

import "fmt"

func binarySearch(arr []int, targetElement int) int {
	left, right := 0, len(arr)-1
	idx := -1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == targetElement {
			idx = mid
			break
		} else if arr[mid] < targetElement {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return idx
}

func main() {
	arr := make([]int, 0, 50)
	// заполняем четными элементами, чтобы было легко наглядно показать
	for i := 0; i <= 100; i += 2 {
		arr = append(arr, i)
	}

	fmt.Println(arr)
	fmt.Println("index of element:", binarySearch(arr, 78))
}
