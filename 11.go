package main

import "fmt"

func convertToMap(arr []int) map[int]struct{} {
	result := make(map[int]struct{})

	for _, v := range arr {
		result[v] = struct{}{}
	}

	return result
}

func intersect(a, b []int) {
	firstSet := convertToMap(a)
	secondSet := convertToMap(b)

	if len(firstSet) > len(secondSet) {
		firstSet, secondSet = secondSet, firstSet
	}

	for k, _ := range firstSet {
		if _, ok := secondSet[k]; ok {
			fmt.Println(k)
		}
	}
}

func main() {
	A := []int{1, 2, 3, 4}
	B := []int{2, 3, 4}

	intersect(A, B)
}
