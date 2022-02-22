package visitor

import "fmt"

type ExportXMLVisitor struct{}

func (visitor *ExportXMLVisitor) visitCircle(c *Circle) {
	fmt.Printf("Exporting XML of Circle with coordinates (%d, %d)\n", c.x, c.y)
}

func (visitor *ExportXMLVisitor) visitRectangle(r *Rectangle) {
	fmt.Printf("Exporting XML of Rectangle with coordinates (%d, %d)\n", r.x, r.y)
}

func (visitor *ExportXMLVisitor) visitDot(d *Dot) {
	fmt.Printf("Exporting XML of Dot with coordinates (%d, %d)\n", d.x, d.y)
}

func NewExportXMLVisitor() *ExportXMLVisitor {
	return &ExportXMLVisitor{}
}
