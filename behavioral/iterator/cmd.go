package iterator

import "fmt"

/*
迭代器是一种行为设计模式，让你能在不暴露复杂数据结构内部细节的情况下遍历其中所有的元素。

在迭代器的帮助下，客户端可以用一个迭代器接口以相似的方式遍历不同集合中的元素。

迭代器模式的主要思想是将集合背后的迭代逻辑提取至不同的、名为迭代器的对象中。
此迭代器提供了一种泛型方法，可用于在集合上进行迭代，而又不受其类型影响。
*/

func execute() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
