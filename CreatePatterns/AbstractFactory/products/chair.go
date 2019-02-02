package products

import "fmt"

type Chair interface {
	HasLegs()
}

//
type ModernChair struct{}

func (*ModernChair) HasLegs() {
	fmt.Println("I am modernChair")
}

//
type ArtDecorChair struct{}

func (*ArtDecorChair) HasLegs() {
	fmt.Println("I am artDecorChair")
}

//
type VictorianChair struct{}

func (*VictorianChair) HasLegs() {
	fmt.Println("I am victorianChair")
}
