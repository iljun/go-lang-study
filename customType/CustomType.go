package main

import "fmt"

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)

	carFuel2 := Gallons(1.2)
	busFuel2 := Liters(4.5)
	carFuel2 += ToGallons(Liters(40.0))
	busFuel2 += ToLiters(Gallons(30.0))
	fmt.Println(carFuel2)
	fmt.Println(busFuel2)
}
