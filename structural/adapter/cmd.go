package adapter

/*
适配器是一种结构型设计模式，它能使不兼容的对象能够相互合作。
适配器可担任两个对象间的封装器，它会接收对于一个对象的调用，
并将其转换为另一个对象可识别的格式和接口。
*/

func execute() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
