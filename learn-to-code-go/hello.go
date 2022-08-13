package main

import "rsc.io/quote"
import "fmt"

var a int = 43
var c int

type test int

var b test

func main() {
	x := 42
	b = 1290
	fmt.Println(x)
	fmt.Println(b)
	b = 1200
	fmt.Println(b)

	if x == 42 {
		foo()
	}
	//type conversion
	a = int(b)
	fmt.Printf("%T\n", a)
}

func foo() {
	fmt.Println("in foo")
}

func Hello() string {
	return quote.Hello()
}
