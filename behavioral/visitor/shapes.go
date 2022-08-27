package visitor

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visit(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visit(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visit(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}
