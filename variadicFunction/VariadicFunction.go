package main

import "fmt"

func main() {

	myFunc(1, "a", "b")
	myFunc(2, "c", "d", "e", "f")
	mySlice := []string{"x", "y", "z"}
	myFunc(3, mySlice...)

}

func myFunc(param1 int, param2 ...string) {
	fmt.Println(param1)
	fmt.Println(param2)
}
