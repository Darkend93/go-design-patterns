package main

import "fmt"

type PizzaType int

type PizzaSize int

const (
	Margherita PizzaType = iota
	Pepperoni
	Diablo
	R25 PizzaSize = 25
	R30 PizzaSize = 30
	R35 PizzaSize = 35
)

type pizza struct {
	Type  PizzaType
	Size  PizzaSize
	Price float64
}

func (p *pizza) getPrice() float64 {
	typeMap := map[PizzaType]float64{
		Margherita: 1700,
		Pepperoni:  2000,
		Diablo:     2200,
	}

	sizeMap := map[PizzaSize]float64{
		R25: 1,
		R30: 1.25,
		R35: 1.5,
	}

	return typeMap[p.Type] * sizeMap[p.Size]
}

func (p *pizza) String() string {
	return fmt.Sprintf("This is %s pizza of size %v. Price: %.2f", p.Type.String(), p.Size, p.getPrice())
}

type Pizza interface {
	getPrice() float64
	String() string
}

func (p PizzaType) String() string {
	return [...]string{"Margherita", "Pepperoni", "Diablo"}[p]
}

func newPizzaFactory(name string, size PizzaSize) Pizza {
	priceMap := make(map[PizzaSize]float64)
	priceMap[R25] = 1
	priceMap[R30] = 1.25
	priceMap[R35] = 1.5

	switch name {
	case Margherita.String():
		return &pizza{
			Type: Margherita,
			Size: size,
		}
	case Pepperoni.String():
		return &pizza{
			Type: Pepperoni,
			Size: size,
		}
	case Diablo.String():
		return &pizza{
			Type: Diablo,
			Size: size,
		}
	default:
		panic("invalid pizza name")
	}

}
