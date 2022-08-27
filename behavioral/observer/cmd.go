package observer

/*
观察者是一种行为设计模式，允许一个对象将其状态的改变通知其他对象。
观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制，可对其事件进行订阅和取消订阅。
*/

func execute() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
