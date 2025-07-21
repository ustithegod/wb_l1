package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v interface{} = make(chan string)

	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("undefined")
		}
	}
}
