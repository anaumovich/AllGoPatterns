package products

import "fmt"

type Cupboard interface {
	HasShelf()
}

//
type ModernCupboard struct{}

func (*ModernCupboard) HasShelf() {
	fmt.Println("I am modernCupboard")
}

//
type ArtDecorCupboard struct{}

func (*ArtDecorCupboard) HasShelf() {
	fmt.Println("I am artDecorCupboard")
}

//
type VictorianCupboard struct{}

func (*VictorianCupboard) HasShelf() {
	fmt.Println("I am victorianCupboard")
}
