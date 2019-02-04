package AbstractFactory

import "AllGoPatterns/CreatePatterns/AbstractFactory/products"

type VictorianFurnitureFactory struct{}

func (*VictorianFurnitureFactory) CreateChair() products.Chair {
	return new(products.VictorianChair)
}

func (*VictorianFurnitureFactory) CreateSofa() products.Sofa {
	return new(products.VictorianSofa)
}

func (*VictorianFurnitureFactory) CreateCupboard() products.Cupboard {
	return new(products.VictorianCupboard)
}
