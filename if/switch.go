package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//i := 3
	i := randNum()
	switch i {
	case 1, 2, 3:
		fmt.Printf("點數: %d  [小]\n", i)
	case 4, 5, 6:
		fmt.Printf("點數: %d  [大]\n", i)
	default:
		fmt.Println("你出老千")
	}

	switch {
	case i > 0 && i < 4:
		fmt.Printf("點數: %d  [小]\n", i)
	case i < 7 && i > 3:
		fmt.Printf("點數: %d  [大]\n", i)
	default:
		fmt.Println("你出老千")
	}

}

func randNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6)
}
