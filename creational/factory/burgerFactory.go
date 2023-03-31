package main

type BurgerType int

type Size float64

func (b Size) String() string {
	switch b {
	case Big:
		return "Big"
	case Double:
		return "Double"
	case Triple:
		return "Triple"
	}
	return ""
}

const (
	Classic BurgerType = iota
	CheeseBurger
	FatJoe
	Big    Size = 1
	Double Size = 1.25
	Triple Size = 1.5
)

func (p BurgerType) String() string {
	return [...]string{"Classic", "CheeseBurger", "FatJoe"}[p]
}

func newBurgerFactory(b BurgerType, size Size) Burger {
	switch b {
	case Classic:
		return NewClassicBurger(size)
	case CheeseBurger:
		return NewCheeseBurger(size)
	default:
		panic("invalid pizza name")
	}

}
