package decorator

import "fmt"

/*
装饰是一种结构设计模式，允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。

由于目标对象和装饰器遵循同一接口，因此你可用装饰来对对象进行无限次的封装。
结果对象将获得所有封装器叠加而来的行为。

与proxy模式有点相似。
*/

func execute() {

	pizza := &VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
