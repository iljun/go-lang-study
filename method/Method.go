package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi")
}

func (m MyType) WithReturn() int {
	return len(m)
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	value := MyType("MyType Value")
	value.sayHi()
	fmt.Println(value.WithReturn())

	number := Number(99)
	fmt.Println("origin value :", number)
	number.Double()
	fmt.Println("double value:", number)
}
