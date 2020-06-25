package main

import "fmt"

type bot struct {
	name string
	age  int
	food string
}

func (p *bot) Name() {
	fmt.Printf("[BOT] My name is %s \n", p.name)
}

func (p *bot) Age() {
	fmt.Printf("[BOT] My age is %d\n", p.age)
}

func (p *bot) Eat() {
	fmt.Printf("[BOT] I eat %s\n", p.food)
}
