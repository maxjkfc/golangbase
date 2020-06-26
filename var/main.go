package main

import "fmt"

func main() {
	//v()
	//v1()
	//v2()
	v3()
}

func v() {
	var a string = "Hello"

	fmt.Printf("%T , %s\n", a, a)

	var b int = 129393020
	fmt.Printf("%T , %d\n", b, b)

	var c float64 = 0.002
	fmt.Printf("%T , %f\n", c, c)

}

func v1() {
	var (
		a string  = "Hello"
		b int     = 129393029
		c float64 = 0.002
	)

	fmt.Printf("%T , %s\n", a, a)
	fmt.Printf("%T , %d\n", b, b)
	fmt.Printf("%T , %f\n", c, c)
}

func v2() {
	a := "Hello"
	b := 129393029
	c := 0.002

	fmt.Printf("%T , %s\n", a, a)
	fmt.Printf("%T , %d\n", b, b)
	fmt.Printf("%T , %f\n", c, c)
}

func v3() {
	a, b, c := "Hello", 129393029, 0.002
	fmt.Printf("%T , %v\n", a, a)
	fmt.Printf("%T , %v\n", b, b)
	fmt.Printf("%T , %v\n", c, c)

	x, y, z := a, c, b
	fmt.Printf("%T , %v\n", x, x)
	fmt.Printf("%T , %v\n", y, y)
	fmt.Printf("%T , %v\n", z, z)
}

