package factory

import "fmt"

/*
工厂方法定义了一个方法，
且必须使用该方法代替通过直接调用构造函数来创建对象 （ new操作符） 的方式。
子类可重写该方法来更改将被创建的对象所属类。
*/

func execute() {
	f := &gunFactory{}
	ak47, _ := f.getGun("ak47")
	musket, _ := f.getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
