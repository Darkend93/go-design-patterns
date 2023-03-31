package main

type Pizza struct {
	Size      PizzaSize
	DoughType DoughType
	Toppings  []Topping
	Price     float64
}

type PizzaSize int

type DoughType int

const (
	Thin DoughType = iota
	Thick
	BasePrice float64   = 1500.00
	R25       PizzaSize = 25
	R30       PizzaSize = 30
	R35       PizzaSize = 35
)

var (
	Olives   Topping = Topping{Name: "Olives", Price: 250.00}
	Salami   Topping = Topping{Name: "Salami", Price: 500.00}
	Tomatoes Topping = Topping{Name: "Tomatoes", Price: 250.00}
	Jalapeno Topping = Topping{Name: "Jalapeno", Price: 350.00}
)

func (d DoughType) String() string {
	return [...]string{"Thin", "Thick"}[d]
}

type Topping struct {
	Name  string
	Price float64
}

type pizzaBuilder struct {
	Size      PizzaSize
	DoughType DoughType
	Toppings  []Topping
}

type PizzaBuilder interface {
	SetSize(s PizzaSize) PizzaBuilder
	SetDoughType(t DoughType) PizzaBuilder
	AddTopping(t Topping) PizzaBuilder
	AddToppings(t []Topping) PizzaBuilder
	Build() Pizza
}

func NewPizzaBuilder() PizzaBuilder {
	return &pizzaBuilder{}
}

func (p *pizzaBuilder) SetSize(s PizzaSize) PizzaBuilder {
	p.Size = s
	return p
}

func (p *pizzaBuilder) SetDoughType(t DoughType) PizzaBuilder {
	p.DoughType = t
	return p
}

func (p *pizzaBuilder) AddTopping(t Topping) PizzaBuilder {
	p.Toppings = append(p.Toppings, t)
	return p
}

func (p *pizzaBuilder) AddToppings(t []Topping) PizzaBuilder {
	p.Toppings = append(p.Toppings, t...)
	return p
}

func (p *pizzaBuilder) Build() Pizza {
	if p.DoughType == 0 {
		panic("dough type is not set")
	}
	if p.Size == 0 {
		panic("pizza size is not set")
	}

	pizza := Pizza{
		Size:      p.Size,
		DoughType: p.DoughType,
		Toppings:  p.Toppings,
		Price:     BasePrice,
	}
	switch p.Size {
	case R25:
		break
	case R30:
		pizza.Price = pizza.Price + 200
	case R35:
		pizza.Price = pizza.Price + 500
	}
	for _, t := range p.Toppings {
		pizza.Price = pizza.Price + t.Price
	}
	return pizza
}
