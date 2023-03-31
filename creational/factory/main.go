package main

import "fmt"

func main() {
	marg := newPizzaFactory("Margherita", R25)
	dia := newPizzaFactory("Diablo", R35)
	// pollo := newPizzaFactory("Pollo", 40)

	fmt.Println(marg)
	fmt.Println(dia)
	// fmt.Println(pollo)
}
