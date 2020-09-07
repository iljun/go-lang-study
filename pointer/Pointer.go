package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	p = newInt1()
	fmt.Println(*p)
	fmt.Println(p)

	p = newInt2()
	fmt.Println(*p)
	fmt.Println(p)
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
