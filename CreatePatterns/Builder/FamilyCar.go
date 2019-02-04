package Builder

//this struct implemet interface buildProcess
type FamilyCarBuilder struct {
	FamilyCar FamilyCar
}

// this struct will be return
type FamilyCar struct {
	engine,
	transmission,
	wheels string
}

func (c *FamilyCarBuilder) AddEngine() buildProcess {
	c.FamilyCar.engine = "2.0 i"
	return c
}

func (c *FamilyCarBuilder) AddTransmission() buildProcess {
	c.FamilyCar.engine = "mono"
	return c
}

func (c *FamilyCarBuilder) AddWheels() buildProcess {
	c.FamilyCar.engine = "4"
	return c
}

func (c *FamilyCarBuilder) GetResult() FamilyCar {

	return c.FamilyCar
}
