package main

import "fmt"

var packageVariable = "package"

func main() {
	step1()
	step2()
}

func step1() {
	var functionVariable = "function"
	if true {
		var conditionalVariable = "conditional"
		fmt.Println(packageVariable)
		fmt.Println(functionVariable)
		fmt.Println(conditionalVariable)
	}
	fmt.Println(packageVariable)
	fmt.Println(functionVariable)
}

func step2() {
	grade := 60
	var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println(status)
}
