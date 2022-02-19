package main

type IVisitor interface {
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
	visitDot(*Dot)
}
