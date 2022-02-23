package main

import visitor "base/BehavioralPatterns/Visitor/files"

func main() {
	// Concrete visitor
	exportXML := visitor.NewExportXMLVisitor()

	// Shapes
	rectangle := visitor.NewRectangle()
	dot := visitor.NewDot()
	circle := visitor.NewCircle()

	// adding some coordinates (not required)
	rectangle.Move(1, 1)
	dot.Move(2, 2)
	circle.Move(3, 3)

	// Exporting xml
	rectangle.Accept(exportXML)
	dot.Accept(exportXML)
	circle.Accept(exportXML)

}
