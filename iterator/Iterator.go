package main

import "fmt"

func main() {
	step1()
	step2()
	step3()
	step4()
	step5()
}

func step1() {
	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}
}

func step2() {
	for x := 1; true; x++ {
		fmt.Println("infinite")
		break
	}
}

func step3() {
	for x := 1; false; x++ {
		fmt.Println("ff")
	}
}

func step4() {
	x := 1
	for x <= 3 {
		fmt.Println(x)
		x++
	}
}

func step5() {
	var x int
	for x = 1; x <= 3; x++ {
		fmt.Println(x)
	}
	fmt.Println(x)
}
