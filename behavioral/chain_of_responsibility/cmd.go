package responsibility

/*
责任链是一种行为设计模式，允许你将请求沿着处理者链进行发送，直到其中一个处理者对其进行处理。

该模式允许多个对象来对请求进行处理，而无需让发送者类与具体接收者类相耦合。
链可在运行时由遵循标准处理者接口的任意处理者动态生成。
*/

func execute() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
