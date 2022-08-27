package flyweight

import "fmt"

/*
享元是一种结构型设计模式，它允许你在消耗少量内存的情况下支持大量对象。
模式通过共享多个对象的部分状态来实现上述功能。
换句话来说，享元会将不同对象的相同数据进行缓存以节省内存。

在享元模式中，我们会将享元对象存储在map容器中。
每当创建共享享元对象的其他对象时，都会从map容器中获取享元对象。

下面让我们来看看此类安排的内部状态和外部状态：
内部状态：内部状态的服装可在多个恐怖分子和反恐精英对象间共享。
外部状态：玩家位置和玩家所使用的武器就是外部状态，因为其在每个对象中都是不同的。
*/

func execute() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
