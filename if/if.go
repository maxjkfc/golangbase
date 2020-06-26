package main

import "fmt"

func main() {

	if true {
		fmt.Println(true)
	}

	if true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	if x := true; x {
		fmt.Println(x)
	} else if y := true; y {

	} else if !x {

	}

}
