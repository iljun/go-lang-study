package main

import "fmt"

func main() {
	var mySlice []string
	mySlice = make([]string, 7)

	fmt.Println(mySlice)

	literal()
	operator()
	internalArray()
	appendFunction()
}

func literal() {
	myInt := []int{9, 8, 11}
	fmt.Println(myInt)
}

func operator() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]

	fmt.Println(underlyingArray)
	fmt.Println(slice1)

	slice2 := underlyingArray[:4]
	fmt.Println(slice2)

	slice3 := underlyingArray[3:]
	fmt.Println(slice3)
}

func internalArray() {
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	fmt.Println(array1)
	fmt.Println(slice1)

	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)
}

func appendFunction() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))
	slice = append(slice, "d")
	fmt.Println(slice, len(slice))
}

