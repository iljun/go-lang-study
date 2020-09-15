package main

import (
	"fmt"
)

type part struct {
	description string
	count		int
}

type car struct {
	name 	string
	topSpeed float64
}

func createCar(name string, topSpeed float64) car {
	var c  car
	c.name = name
	c.topSpeed = topSpeed
	return c
}

func copyAndPaste(c *car) {
	c.name = c.name + c.name
}

type subscriber struct {
	name string
	rate float64
	active bool
	homeAddress Address
}

type Address struct {
	City string
}

type subscriber2 struct {
	name string
	rate float64
	active bool
	Address
}

func main() {
	var myStruct struct {
		number int
		work string
		toggle bool
	}
	fmt.Println(myStruct)

	myStruct.work = "software developer"
	fmt.Println(myStruct)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println(porsche)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println(bolts)

	c := createCar("BMW", 98123)
	fmt.Println(c)

	copyAndPaste(&c)
	fmt.Println(c)

	var literal car = car{name: "K7", topSpeed: 1234}
	fmt.Println(literal)

	var literal2 car = car{name: "K3"}
	fmt.Println(literal2)

	address := Address{City: "Seoul"}
	subscriber := subscriber{name: "test", rate: 123, active: false, homeAddress: address}
	fmt.Println(subscriber)
	fmt.Println(subscriber.homeAddress)

	subscriber2 := subscriber2{name: "subscriber2", rate: 234, active: false}
	subscriber2.City = "Busan"
	fmt.Println(subscriber2)
}
