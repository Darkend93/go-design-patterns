package main

import "fmt"

type burger struct {
	Size  Size
	Price float64
}

type Burger interface {
	getPrice() float64
	String() string
}
type classicBurger struct {
	burger
}

func (c classicBurger) getPrice() float64 {
	return c.Price
}

func (c classicBurger) String() string {
	return fmt.Sprintf("This is Classic Burger. Price: %.2f", c.getPrice())
}

func NewClassicBurger(size Size) Burger {
	return &classicBurger{
		burger: burger{
			Size:  size,
			Price: 1800 * float64(size),
		},
	}
}

type cheeseBurger struct {
	burger
}

func (c cheeseBurger) getPrice() float64 {
	return c.Price
}

func (c cheeseBurger) String() string {
	return fmt.Sprintf("This is Cheese Burger. Price: %.2f", c.getPrice())
}

func NewCheeseBurger(size Size) Burger {
	return &cheeseBurger{
		burger: burger{
			Size:  size,
			Price: 2000 * float64(size),
		},
	}
}
