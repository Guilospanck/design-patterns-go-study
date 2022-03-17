package implementations

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) SetSize(size int) {
	s.size = size
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

func (s *Shirt) GetSize() int {
	return s.size
}

func NewShirt() *Shirt {
	return &Shirt{}
}
