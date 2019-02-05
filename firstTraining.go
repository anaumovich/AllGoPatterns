package ggg

import (
	"AllGoPatterns/SOLID"
	"fmt"
)

func main() {
	Animal := SOLID.NewAnimal(false, 4, "Animal")
	Cat := SOLID.NewCat(*Animal)
	BengalCat := SOLID.NewBengalCat(*Cat)

	BengalCat.Feed()

	BengalCat.Feed(Cat)
	//Animal.Feed(Cat)
	_, _ = fmt.Println(Animal)
	fmt.Println()
}
