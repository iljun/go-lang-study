package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

	literal()
	length()
	forRange()
}

func literal() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes)

	text := [3]string {
		"This is a series of long strings",
		"which would be awkward of place",
		"together on a single line",
	}

	fmt.Println(text)
}

func length() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))
}

func forRange() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for index, value := range notes {
		fmt.Println(index, value)
	}

	for _, value := range notes {
		fmt.Println(value)
	}

	for index, _ := range notes {
		fmt.Println(index)
	}
}