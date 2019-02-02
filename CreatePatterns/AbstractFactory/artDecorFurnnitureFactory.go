package AbstractFactory

import "AllGoPatterns/CreatePatterns/AbstractFactory/products"

type ArtDecorFurnitureFactory struct {
}

func NewArtDecorFurnitureFactory() *ArtDecorFurnitureFactory {
	return &ArtDecorFurnitureFactory{}
}

func (*ArtDecorFurnitureFactory) CreateChair() products.Chair {
	return new(products.ArtDecorChair)
}

func (*ArtDecorFurnitureFactory) CreateSofa() products.Sofa {
	return new(products.ArtDecorSofa)
}

func (*ArtDecorFurnitureFactory) CreateCupboard() products.Cupboard {
	return new(products.ArtDecorCupboard)
}
