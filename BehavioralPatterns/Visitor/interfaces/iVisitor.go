package interfaces

import "base/BehavioralPatterns/Visitor/visitor"

type IVisitor interface {
	visitCircle(*visitor.Circle)
	visitRectangle(*visitor.Rectangle)
	visitDot(*visitor.Dot)
}
