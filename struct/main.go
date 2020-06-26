package main

import (
	"fmt"
	"time"
)

type People struct {
	Name     string
	Age      int
	Birthday time.Time
	Food
}

type Food struct {
	Name string
}

func (f Food) Eat() {
	fmt.Printf("Eat Food: %v\n", f.Name)
}

func (p *People) Introduction() {
	fmt.Printf("My Name is :%v . I'm %d years old . My Birthday is : %s\n", p.Name, p.Age, p.Birthday.Format("2006-01-02"))
}

func (p *People) String() string {
	return fmt.Sprintf("Name: %v  Age: %v Birthday: %v", p.Name, p.Age, p.Birthday)
}

func main() {

	max := People{
		Name:     "Max",
		Age:      30,
		Birthday: time.Date(1989, time.February, 9, 0, 0, 0, 0, time.Local),
		Food:     Food{Name: "漢堡"},
	}

	max.Introduction()
	max.Eat()

	//fmt.Printf("%+v\n", max)
	//fmt.Println(max.String())

}
