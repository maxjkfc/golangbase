package main

import (
	"fmt"
	"time"
)

var i int64

func main() {
	i = 0
	single()
	i = 0
	multi()
}
func single() {

	t := time.Now()
	for i < 10000000 {
		i++
		fmt.Println(i)
	}

	fmt.Printf("single total :%v\n cost: %v\n", i, time.Now().Sub(t))
}

func multi() {
	t := time.Now()
	for i < 10000000 {
		go func() {
			i++
			fmt.Println(i)
		}()
	}
	fmt.Printf("multi total :%v\n cost: %v\n", i, time.Now().Sub(t))
}
