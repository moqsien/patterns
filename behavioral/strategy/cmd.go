package strategy

/*
策略是一种行为设计模式，它将一组行为转换为对象，并使其在原始上下文对象内部能够相互替换。

原始对象被称为上下文，它包含指向策略对象的引用并将执行行为的任务分派给策略对象。
为了改变上下文完成其工作的方式，其他对象可以使用另一个对象来替换当前链接的策略对象。

思考一下构建内存缓存的情形。由于处在内存中，故其大小会存在限制。
在达到其上限后，一些条目就必须被移除以留出空间。
如何将缓存类与这些过期移除算法解耦，以便在运行时更改算法。
此外，在添加新算法时，缓存类不应改变。这就是策略模式发挥作用的场景。
*/

func execute() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
