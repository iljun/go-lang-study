package main

import "fmt"

func main() {
	var ranks map[string]int
	ranks = make(map[string]int)

	fmt.Println(ranks)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	fmt.Println(ranks)
	fmt.Println(ranks["TT"])

	literal()
	emptyMap()
	valid()
	deleteFunction()
	rangeFunction()
}

func literal() {
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap)

	myMap2 := map[string]string {
		"H": "Hydrogen",
		"Li": "Lithium",
	}
	fmt.Println(myMap2)
}

func emptyMap() {
	emptyMap := map[string]float64{}

	fmt.Println(emptyMap)
}

func valid() {
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
}

func deleteFunction() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
}

func rangeFunction() {
	grades := map[string]int{"Alma": 74, "Rohit": 85, "Carl": 59}
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %d\n", name, grade)
	}
}
