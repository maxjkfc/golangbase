package main

import "fmt"

func main() {
	var people map[int]string

	people = make(map[int]string, 3)

	people[0] = "tom"
	people[1] = "jam"
	people[2] = "ajx"

	fmt.Println(people[1])

	var peopleLove map[string]string
	peopleLove = make(map[string]string, 3)

	peopleLove["tom"] = "marry"
	peopleLove["jam"] = "mia"
	peopleLove["ajx"] = "java"

	fmt.Println(peopleLove["tom"])

	randMan := map[int]string{
		1: "max",
		2: "herry",
		3: "hardy",
		4: "sylvia",
		5: "thomas",
		6: "linda",
	}

	for _, v := range randMan {
		fmt.Println(v)
	}

	delete(randMan, 1)
	for _, v := range randMan {
		fmt.Println(v)
	}

}
