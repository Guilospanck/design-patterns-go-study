package visitor

var VisitCircleCalled bool = false
var VisitRectangleCalled bool = false
var VisitDotCalled bool = false

type ExportXMLVisitorMock struct{}

func (mock *ExportXMLVisitorMock) visitCircle(v *Circle) {
	VisitCircleCalled = true
}

func (mock *ExportXMLVisitorMock) visitRectangle(v *Rectangle) {
	VisitRectangleCalled = true
}

func (mock *ExportXMLVisitorMock) visitDot(v *Dot) {
	VisitDotCalled = true
}

func NewExportXMLVisitorMock() *ExportXMLVisitorMock {
	return &ExportXMLVisitorMock{}
}
