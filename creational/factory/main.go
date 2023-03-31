package main

import "fmt"

func main() {
	fj := newBurgerFactory(FatJoe, Big)
	cb := newBurgerFactory(CheeseBurger, Double)
	// pollo := newBurgerFactory("Pollo", 40)

	fmt.Println(fj)
	fmt.Println(cb)
	// fmt.Println(pollo)
}
