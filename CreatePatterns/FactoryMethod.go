package CreatePatterns

//FactoryMethod
// What I can say about this pattern.
// Данный паттерн предлагает произвести замену  конкретного
// поля структуры-родителя на метод, который будет выполнять подстановку структуры-ребенка
// необходимой для конкретной реализации. Т.е. для базовой структуры создаем метод-селектор
// который будет выбирать необходимую реализацию методов описанных интерфейсом. А у этого интерфейса
// может быть много структур реализующих требуемый метод.
// Что нам это дает? По сути вместо конкретной структуры мы завязываемся на метод который будет
// возвращать те или иные структуры в зависимости от пердаваемых параметров этому методу, соответствено
// это позволит повысить гибкость кода и позволит добавить другие варианты бизнес логики
// например способы оплаты или транспорт для доставки товара, избежав при этом перелопачивания всей программы

import "fmt"

type logistic struct {
	planDelivery string
	transport    transport
}

func (*logistic) transportCreator(x string) transport {
	switch x {
	case "plane":
		return new(airplane)
	case "car":
		return new(car)
	}
	return nil
}

type transport interface {
	deliver()
}

type airplane struct{}

func (*airplane) deliver() {
	fmt.Println("I am plane")
}

type car struct{}

func (*car) deliver() {
	fmt.Println("I am a car")
}

func main() {
	logistic := new(logistic)
	logistic.transport = logistic.transportCreator("car")
	logistic.transport.deliver()
	logistic.transportCreator("plane").deliver()
}
