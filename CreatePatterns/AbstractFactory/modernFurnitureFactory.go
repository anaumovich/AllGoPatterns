package AbstractFactory

import "AllGoPatterns/CreatePatterns/AbstractFactory/products"

type ModernFurnitureFactory struct{}

func (*ModernFurnitureFactory) CreateChair() products.Chair {
	return new(products.ModernChair)
}

func (*ModernFurnitureFactory) CreateSofa() products.Sofa {
	return new(products.ModernSofa)
}

func (*ModernFurnitureFactory) CreateCupboard() products.Cupboard {
	return new(products.ModernCupboard)
}
