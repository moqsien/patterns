package iterator

type User struct {
	name string
	age  int
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type Collection interface {
	createIterator() Iterator
}
