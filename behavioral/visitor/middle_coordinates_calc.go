package visitor

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visit(s Shape) {
	switch s.(type) {
	case *Square:
		fmt.Println("Calculating middle point coordinates for square")
	case *Circle:
		fmt.Println("Calculating middle point coordinates for circle")
	case *Rectangle:
		fmt.Println("Calculating middle point coordinates for rectangle")
	default:
		fmt.Println("err")
	}
}
