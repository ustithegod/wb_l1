package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Salary float64
}

func (h Human) Greet() string {
	return fmt.Sprintf("Hello, my name is %s", h.Name)
}

func (h Human) CalculateAnnual() float64 {
	return h.Salary * 12
}

type Action struct {
	Human
}

func main() {
	action := Action{Human: Human{Name: "Saveliy", Age: 19, Salary: 10}}
	fmt.Println(action.Greet())
	fmt.Println(action.CalculateAnnual())
}
