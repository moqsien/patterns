package state

import (
	"fmt"
	"log"
)

/*
状态是一种行为设计模式，让你能在一个对象的内部状态变化时改变其行为。
该模式将与状态相关的行为抽取到独立的状态类中，让原对象将工作委派给这些类的实例，
而不是自行进行处理。

当对象可以处于许多不同的状态中时应使用状态设计模式，
同时根据传入请求的不同，对象需要变更其当前状态。
在例子中，自动售货机可以有多种不同的状态，同时会在这些状态之间持续不断地互相转换。

自动售货机的代码不会被状态逻辑行为污染； 所有依赖于状态的代码都存在于各自的状态实现中。
*/
func execute() {
	vendingMachine := newVendingMachine(1, 10)

	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
