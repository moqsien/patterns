package mediator

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
