package Builder

type SportCarBuilder struct {
	SportCar SportCar
}

func newSportCarBuilder(sportCar SportCar) *SportCarBuilder {
	return &SportCarBuilder{SportCar: sportCar}
}

type SportCar struct {
	engine,
	transmission,
	wheels string
}

func newSportCar(engine string, transmission string, wheels string) *SportCar {
	return &SportCar{engine: engine, transmission: transmission, wheels: wheels}
}

func (c *SportCarBuilder) AddEngine() buildProcess {
	c.SportCar.engine = "5.0 ti"
	return c
}

func (c *SportCarBuilder) AddTransmission() buildProcess {
	c.SportCar.transmission = "full"
	return c
}

func (c *SportCarBuilder) AddWheels() buildProcess {
	c.SportCar.wheels = "4"
	return c
}

func (c *SportCarBuilder) getResult() SportCar {

	return c.SportCar
}
