package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b int64
	fmt.Print("enter a: ")
	fmt.Scanln(&a)
	fmt.Print("enter b: ")
	fmt.Scanln(&b)

	bigA := big.NewInt(a)
	bigB := big.NewInt(b)

	var bigOp big.Int

	fmt.Println("сложение: ", bigOp.Add(bigA, bigB))
	fmt.Println("вычитание: ", bigOp.Sub(bigA, bigB))
	fmt.Println("умножение: ", bigOp.Mul(bigA, bigB))
	fmt.Println("деление: ", bigOp.Div(bigA, bigB))

}
