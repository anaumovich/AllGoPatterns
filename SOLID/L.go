package SOLID

// This File describes base principles of Barbara Liskov (
//a series of compatibility checks child class with the parent class ).
// First principle: all parameters of the methods of child class must be compatible
// with parameters of the methods of parent class, or the parameters of the methods of the child
// class must be more abstract, them  parameters of the methods of the parents class
// т.е. Если мы в дочернем классе переопределяем какой-либо метод родительского класса,
// то этот метод должен принимать те же, или более абстрактные аргументы(параметры) что и родительский класс
// Второй принцип подстановок барабары Лисков: Если мы переопределяем один из методов родительского класса
// в классе ребенка, то он должен возвращать тот же тип значения что и родительский метод,
// или возвращаемое значение должно быть подтипом значения, которое возвращает метод класса родителя
// На сколько я понимаю в го как такового наследования нет и вся работа по этому построена через интерфейсы,
// соответственно на этапе сборки проекта компиляция не выполнится или программа будет работать не корректно,
// так как если метод ребенка не будет принимать и возвращать значения интерфейса через который осуществляется
// взаимодействие ребенка и родителя то она не будет считаться работающей т.е. она не реализует интерфейс
// Третий принцип: метод наследника не должен ослаблять или наооборот ужесточать условия с которыми работает метод родителя
// Четвертый принцип Меетод наследника не должен изменять базовые поля которые характеризуют родителя, т.е. кот всегда должен мурчать
// Пятый принцип Ошибки которые возвращает метод класса ребенка должны быть

type Animal struct { //main class
	HasEat     bool
	NumberLegs int
	Name       string
	Cat        hasChild // здесь будет правильным решением использовать интерфей, так и сделаем
}

type hasChild interface {
	someFunc(int)
}

//Correct
type BengalCat struct{}

func (*BengalCat) someFunc(int) {}

//
//Not correct
type BritishCat struct{}

func (*BritishCat) someFunc() {}
