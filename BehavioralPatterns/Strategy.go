package Strategy

var i = "+"

type Order struct {
	id       int
	date     string
	shipping ShippingType
}

type ShippingType interface {
	deliveryPrice()
	GetDate(order Order)
}

type AirDelivery struct{}

func NewAirDelivery() *AirDelivery {
	return &AirDelivery{}
}

func (*AirDelivery) deliveryPrice() {
	//implementation method for AirDelivery
}
func (*AirDelivery) GetDate(order Order) {
	//implementation method for AirDelivery
}

type SeaDelivery struct{}

func NewSeaDelivery() *SeaDelivery {
	return &SeaDelivery{}
}
func (*SeaDelivery) deliveryPrice() {
	//implementation method for SeaDelivery
}
func (*SeaDelivery) GetDate(order Order) {
	//implementation method for SeaDelivery
}

func Controller(i string, order Order) {
	if i == "1" {
		order.shipping = NewAirDelivery()
	}
	if i == "2" {
		order.shipping = NewSeaDelivery()
	}
}
