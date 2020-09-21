package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")

	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<- myChannel)
}
