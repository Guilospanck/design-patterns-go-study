package visitor

type IVisitor interface {
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
	visitDot(*Dot)
}
