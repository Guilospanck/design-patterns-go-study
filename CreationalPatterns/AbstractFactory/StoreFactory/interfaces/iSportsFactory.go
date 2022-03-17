package interfaces

// Abstract Factory interface
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
