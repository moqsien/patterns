package mediator

/*
中介者是一种行为设计模式，让程序组件通过特殊的中介者对象进行间接沟通，
达到减少组件之间依赖关系的目的。

中介者能使得程序更易于修改和扩展，而且能更方便地对独立的组件进行复用，
因为它们不再依赖于很多其他的类。

中介者模式的一个绝佳例子就是火车站交通系统。
两列火车互相之间从来不会就站台的空闲状态而进行通信。
*/
func execute() {
	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
