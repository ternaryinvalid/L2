package main

import "fmt"

type Pizza struct {
	size      string
	cheese    bool
	pepperoni bool
	mushrooms bool
}

type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	Build() Pizza
}

type concretePizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() PizzaBuilder {
	return &concretePizzaBuilder{}
}

func (b *concretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.size = size
	return b
}

func (b *concretePizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.cheese = true
	return b
}

func (b *concretePizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.pepperoni = true
	return b
}

func (b *concretePizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.mushrooms = true
	return b
}

func (b *concretePizzaBuilder) Build() Pizza {
	return b.pizza
}

func main() {
	builder := NewPizzaBuilder()

	pizza := builder.
		SetSize("large").
		AddCheese().
		AddPepperoni().
		Build()

	fmt.Printf("Pizza size: %s, Cheese: %t, Pepperoni: %t, Mushrooms: %t\n",
		pizza.size, pizza.cheese, pizza.pepperoni, pizza.mushrooms)
}
