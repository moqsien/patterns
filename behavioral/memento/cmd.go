package memento

import "fmt"

/*
备忘录是一种行为设计模式，允许生成对象状态的快照并在以后将其还原。
备忘录不会影响它所处理的对象的内部结构，也不会影响快照中保存的数据。

备忘录模式让我们可以保存对象状态的快照。
你可使用这些快照来将对象恢复到之前的状态。这在需要在对象上实现"撤销-重做"操作时非常实用。
*/

func execute() {

	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	// 创建并添加备忘录
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	// 获取备忘录并还原
	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())

}
