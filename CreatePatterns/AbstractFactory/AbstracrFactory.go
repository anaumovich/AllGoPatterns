package AbstractFactory

import (
	"AllGoPatterns/CreatePatterns/AbstractFactory/products"
)

// AbstractFactory: this pattern should creates different product families

type Furniture struct {
	color            string
	furnitureFactory AbstractFurnitureFactory
}

type AbstractFurnitureFactory interface {
	CreateChair() products.Chair
	CreateSofa() products.Sofa
	CreateCupboard() products.Cupboard
}

func CreateFurniture(x string) AbstractFurnitureFactory {
	switch x {
	case "modern":
		return NewModernFurnitureFactory()
	case "artDecor":
		return NewArtDecorFurnitureFactory()
	case "victorian":
		return NewVictorianFurnitureFactory()
	}
	return nil
}

func main() {
	furniture := new(Furniture)
	furniture.furnitureFactory = CreateFurniture("modern")
	furniture.furnitureFactory.CreateCupboard().HasShelf()
	furniture.furnitureFactory.CreateSofa().HasSeat()
	furniture.furnitureFactory.CreateChair().HasLegs()

}
