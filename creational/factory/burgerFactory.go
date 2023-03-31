package main

import "fmt"

type BurgerType int

type BurgerSize float64

func (b BurgerSize) String() string {
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
	Big    BurgerSize = 1
	Double BurgerSize = 1.25
	Triple BurgerSize = 1.5
)

type burger struct {
	Type  BurgerType
	Size  BurgerSize
	Price float64
}

func (p *burger) getPrice() float64 {
	typeMap := map[BurgerType]float64{
		Classic:      1700,
		CheeseBurger: 2000,
		FatJoe:       2200,
	}

	return typeMap[p.Type] * float64(p.Size)
}

func (p *burger) String() string {
	return fmt.Sprintf("This is %s burger of %v size. Price: %.2f", p.Type.String(), p.Size, p.getPrice())
}

type Burger interface {
	getPrice() float64
	String() string
}

func (p BurgerType) String() string {
	return [...]string{"Classic", "CheeseBurger", "FatJoe"}[p]
}

func newBurgerFactory(b BurgerType, size BurgerSize) Burger {
	switch b {
	case Classic:
		return &burger{
			Type: Classic,
			Size: size,
		}
	case CheeseBurger:
		return &burger{
			Type: CheeseBurger,
			Size: size,
		}
	case FatJoe:
		return &burger{
			Type: FatJoe,
			Size: size,
		}
	default:
		panic("invalid pizza name")
	}

}
