package main

import "fmt"

type People interface {
	Name()
	Age()
	Eat()
}

type people struct {
	name string
	age  int
	food string
}

func (p *people) Name() {
	fmt.Printf("[PEOPLE] My name is %s \n", p.name)
}

func (p *people) Age() {
	fmt.Printf("[PEOPLE] My age is %d\n", p.age)
}

func (p *people) Eat() {
	fmt.Printf("[PEOPLE] I eat %s\n", p.food)
}

//
func main() {
	tom := &people{
		name: "tommy",
		age:  10,
		food: "漢堡",
	}

	Room(tom)

	//x1 := &bot{
	//name: "x1",
	//age:  199999999,
	//food: "機油",
	//}
	//Room(x1)

}

func Room(p People) {

	fmt.Println("Some People In Room")
	p.Name()
	p.Age()
	p.Eat()
}
