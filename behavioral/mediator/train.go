package mediator

type Train interface {
	arrive()
	depart()
	permitArrival()
}
