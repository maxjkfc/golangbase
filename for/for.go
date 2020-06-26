package main

import "fmt"

func main() {
	for1()
	for2()
}

func for1() {
	var sum int

	for i := 0; i < 1000; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func for2() {

	var sum, i int

	for i < 1000 {
		sum += i
		i += 1
	}

	fmt.Println(sum)

}

func for3() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range a {
		fmt.Println(i, v)
	}

}
