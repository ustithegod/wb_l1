package main

import "fmt"

func reverseString(s string) string {
	runeSlice := []rune(s)
	i, j := 0, len(runeSlice)-1
	for k := 0; k <= j/2+1; k++ {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
		i++
		j--
	}

	return string(runeSlice)
}

func main() {
	s := "главрыба"
	fmt.Println(reverseString(s))
}
