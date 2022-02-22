package mocks

import "base/BehavioralPatterns/Visitor/visitor"

var VisitCircleCalled bool = false

type ExportXMLVisitorMock struct{}

func (mock *ExportXMLVisitorMock) visitCircle(v *visitor.Circle) {
	VisitCircleCalled = true
}

func (mock *ExportXMLVisitorMock) visitRectangle(v *visitor.Rectangle) {

}

func (mock *ExportXMLVisitorMock) visitDot(v *visitor.Dot) {

}

func NewExportXMLVisitorMock() *ExportXMLVisitorMock {
	return &ExportXMLVisitorMock{}
}
