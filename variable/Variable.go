package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("allocateVariable")
	allocateVariable()

	fmt.Println("initializeVariable")
	initializeVariable()

	fmt.Println("zeroValue")
	zeroValue()

	fmt.Println("shortVariableDeclaration")
	shortVariableDeclaration()

	fmt.Println("conversion")
	conversion()
}

func allocateVariable() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square maters")
	fmt.Println()

}

func initializeVariable() {
	var quantity int = 4
	var length, width float64 = 1.2, 2.4
	var customerName string = "Damon Cole"

	fmt.Println(quantity)
	fmt.Println(length, width)
	fmt.Println(customerName)
	fmt.Println()

}

func zeroValue() {
	var quantity int
	var length, width float64
	var customerName string
	var myBool bool

	// zero value
	// zero value is default value
	fmt.Println(quantity)
	fmt.Println(length, width)
	fmt.Println(customerName)
	fmt.Println(myBool)
}

func shortVariableDeclaration() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"
	myBool := false

	fmt.Println(quantity)
	fmt.Println(length, width)
	fmt.Println(customerName)
	fmt.Println(myBool)
}

func conversion() {
	var quantity int = 2
	var length, width float64 = 1.2, 2.4

	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(float64(quantity)))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(int64(width)))
}