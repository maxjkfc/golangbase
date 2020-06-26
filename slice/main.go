package main

import "fmt"

func main() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("%T  \tcap:%d \tlen:%d \t%v \n", a, cap(a), len(a), a)

	b := make([]int, 0, 0)
	fmt.Printf("%T  \t\tcap:%d \tlen:%d \t%v \n", b, cap(b), len(b), b)

	b = append(b, 101)
	fmt.Printf("%T  \t\tcap:%d \tlen:%d \t%v \n", b, cap(b), len(b), b)

	b = a[0:3]
	fmt.Printf("%T  \t\tcap:%d \tlen:%d \t%v \n", b, cap(b), len(b), b)
}
