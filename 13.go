package main

import "fmt"

func main() {

	a, b := 10, 5

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)
}
