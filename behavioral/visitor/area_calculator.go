package visitor

import (
	"fmt"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visit(s Shape) {
	switch s.(type) {
	case *Square:
		fmt.Println("Calculating area for square")
	case *Circle:
		fmt.Println("Calculating area for circle")
	case *Rectangle:
		fmt.Println("Calculating area for rectangle")
	default:
		fmt.Println("err")
	}
}
