package main

import (
	"fmt"
	"strings"
)

func CheckUnique(s string) bool {
	letters := make(map[string]struct{})
	for i := 0; i < len(s); i++ {
		k := strings.ToLower(string(s[i]))
		if _, exists := letters[k]; exists {
			return false
		}
		letters[k] = struct{}{}
	}

	return true
}

func main() {
	s := "abcd"
	fmt.Printf("\"%s\": %t\n", s, CheckUnique(s))

	s2 := "abCdefA"
	fmt.Printf("\"%s\": %t\n", s2, CheckUnique(s2))

	s3 := "aabcdd"
	fmt.Printf("\"%s\": %t\n", s3, CheckUnique(s3))
}
