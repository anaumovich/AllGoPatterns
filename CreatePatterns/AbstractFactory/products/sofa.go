package products

import "fmt"

type Sofa interface {
	HasSeat()
}

//
type ModernSofa struct{}

func (*ModernSofa) HasSeat() {
	fmt.Println("I am modernSofa")
}

//
type ArtDecorSofa struct{}

func (*ArtDecorSofa) HasSeat() {
	fmt.Println("I am artDecorSofa")
}

//
type VictorianSofa struct{}

func (*VictorianSofa) HasSeat() {
	fmt.Println("I am victorianSofa")
}
