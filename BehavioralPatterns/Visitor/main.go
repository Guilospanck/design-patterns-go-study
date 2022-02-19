package main

func main() {
	// Concrete visitor
	exportXML := NewExportXMLVisitor()

	// Shapes
	rectangle := NewRectangle()
	dot := NewDot()
	circle := NewCircle()

	// adding some coordinates (not required)
	rectangle.Move(1, 1)
	dot.Move(2, 2)
	circle.Move(3, 3)

	// Exporting xml
	rectangle.Accept(exportXML)
	dot.Accept(exportXML)
	circle.Accept(exportXML)

}
