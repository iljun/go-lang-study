package main

import "fmt"

type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(f float64)
	methodWithReturnValue() float64
}

type MyType int

func (m MyType) methodWithoutParameters() {
	fmt.Println("methodWithoutParameters")
}

func (m MyType) methodWithParameter(f float64) {
	fmt.Println("methodWithParameter")
}

func (m MyType) methodWithReturnValue() float64 {
	fmt.Println("methodWithReturnValue")
	return 1.1
}

type myInterface2 interface {
	onlyUseInterface()
}

type MyType2 int

func (m MyType2) onlyUseInterface() {
	fmt.Println("only use Interface")
}

func (m MyType2) typeAssertion() {
	fmt.Println("typeAssertion")
}

type MyType3 int

func (m MyType3) onlyUseInterface() {
	fmt.Println("only use Interface")
}

func (m MyType3) typeAssertion() {
	fmt.Println("typeAssertion")
}

func main() {
	var m myInterface = MyType(1)
	fmt.Println(m)
	m.methodWithoutParameters()
	m.methodWithParameter(1.1)
	fmt.Println(m.methodWithReturnValue())

	var m2 myInterface2 = MyType2(2)
	fmt.Println(m2)
	m2.onlyUseInterface()
	// compile error
	// m2.typeAssertion()

	// type Assertion
	var m3 myInterface2 = MyType3(3)
	var assertion = m3.(MyType3)
	assertion.onlyUseInterface()
	assertion.typeAssertion()
}
