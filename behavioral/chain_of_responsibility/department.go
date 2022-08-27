package responsibility

type Department interface {
	execute(*Patient)
	setNext(Department)
}
