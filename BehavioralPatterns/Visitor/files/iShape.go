package visitor

type IShape interface {
	Move(x, y int)
	Draw()
	Accept(v IVisitor)
}
