package main

import "fmt"

func main() {
	sayHi()
	repeatLine("repeat", 5)
	fmt.Println(calculatePlus(3, 5))
	fmt.Println(multiply(3, 5))
	fmt.Println(multipleReturn())
	myInt, myBool := multipleReturn()
	fmt.Println(myInt, myBool)

	nextInt, nextBool := multipleReturnName()
	fmt.Println(nextInt, nextBool)

	i, err := returnError()
	fmt.Println(i)
	fmt.Println(err)

	j, n := returnNil()
	fmt.Println(j)
	fmt.Println(n)

	p := createPointer()
	fmt.Println(p)
	fmt.Println(*p)

	pi := 3
	printPointer(&pi)
}

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func calculatePlus(num1 int, num2 int) int {
	return num1 + num2
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func multipleReturn() (int, bool) {
	return 3, true
}

func multipleReturnName() (num int, myBool bool) {
	return 4, false
}

func returnError() (num int, err error) {
	return 0, fmt.Errorf("error")
}

func returnNil() (num int, err error) {
	return 1, nil
}

func createPointer() *int {
	pointer := 1
	return &pointer
}

func printPointer(myPointer *int) {
	fmt.Println(*myPointer)
}
