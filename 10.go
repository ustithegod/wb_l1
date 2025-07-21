package main

import (
	"fmt"
	"strconv"
)

func findRange(floatNumber float64) int {
	intNumber := int(floatNumber)
	return intNumber / 10 * 10
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 5.6, -2.3}
	groups := make(map[string][]float64)

	var (
		key   string
		index int
	)

	for _, v := range temperatures {
		index = findRange(v)
		if index == 0 && v >= 0 {
			key = "+0"
		} else if index == 0 && v < 0 {
			key = "-0"
		} else {
			key = strconv.Itoa(index)
		}

		groups[key] = append(groups[key], v)
	}

	fmt.Println(groups)
}
