package main

import "fmt"

func Unify(data []string) []string {
	uniqueMap := make(map[string]struct{})

	for _, v := range data {
		uniqueMap[v] = struct{}{}
	}

	keys := make([]string, 0, len(uniqueMap))
	for k := range uniqueMap {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(Unify(arr))
}
