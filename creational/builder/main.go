package main

import "fmt"

func main() {
	pizza1 := NewPizzaBuilder().
		SetSize(R30).
		SetDoughType(Thin).
		AddTopping(Olives).
		AddTopping(Tomatoes).
		Build()

	pizza2 := NewPizzaBuilder().
		AddToppings([]Topping{Olives, Tomatoes, Jalapeno, Salami}).
		SetSize(R35).
		SetDoughType(Thick).
		Build()

	//pizza3 := NewPizzaBuilder().Build()

	fmt.Println(pizza1)
	fmt.Println(pizza2)
	//fmt.Println(pizza3)
}
