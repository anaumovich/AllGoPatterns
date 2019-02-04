package Builder

import "fmt"

type buildProcess interface {
	AddEngine() buildProcess
	AddTransmission() buildProcess
	AddWheels() buildProcess
}

type ManufacturingDirector struct {
	builder buildProcess
}

func (f *ManufacturingDirector) SetBuilder(b buildProcess) {
	f.builder = b
}
func (f *ManufacturingDirector) Construct() {
	f.builder.AddWheels().AddEngine().AddTransmission()
}

func main() {

	director := new(ManufacturingDirector)
	car := newSportCar("", "", "")
	carBuilder := newSportCarBuilder(*car)
	director.SetBuilder(carBuilder)
	director.Construct()
	fmt.Println(director.builder)
}
