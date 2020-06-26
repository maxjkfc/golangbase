package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T %v\n", a, a)
	a[0] = 100
	fmt.Printf("%T %v\n", a, a)

	//a = append(a, 101)
	//fmt.Printf("%T %v\n", a, a)

	fmt.Println(a[0:3])
	fmt.Println(a[3:len(a)])

	b := a
	fmt.Printf("%T %v\n", b, b)
	c := a[0:3]
	fmt.Printf("%T %v\n", c, c)
}
