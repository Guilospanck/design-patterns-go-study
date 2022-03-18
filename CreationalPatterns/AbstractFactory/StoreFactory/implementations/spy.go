package implementations

import "github.com/stretchr/testify/mock"

type shoeSpy struct {
	mock.Mock
	Size int
	Logo string
}

func (spy *shoeSpy) SetLogo(logo string) {
	spy.Called(logo)
}

func (spy *shoeSpy) SetSize(size int) {
	spy.Called(size)
}

func (spy *shoeSpy) GetSize() int {
	spy.Called()
	return spy.Size
}

func (spy *shoeSpy) GetLogo() string {
	spy.Called()
	return spy.Logo
}

// func NewShoeSpy() *shoeSpy {
// 	Size := 10
// 	Logo := "logo"

// 	return &shoeSpy{Size: Size, Logo: Logo}
// }

type adidasShoeSpy struct {
	shoeSpy
}

type nikeShoeSpy struct {
	shoeSpy
}

//

type shirtSpy struct {
	mock.Mock
	Size int
	Logo string
}

func (spy *shirtSpy) SetLogo(logo string) {
	spy.Called(logo)
}

func (spy *shirtSpy) SetSize(size int) {
	spy.Called(size)
}

func (spy *shirtSpy) GetSize() int {
	spy.Called()
	return spy.Size
}

func (spy *shirtSpy) GetLogo() string {
	spy.Called()
	return spy.Logo
}

// func NewShirtSpy() *shirtSpy {
// 	Size := 10
// 	Logo := "logo"

// 	return &shirtSpy{Size: Size, Logo: Logo}
// }

type adidasShirtSpy struct {
	shirtSpy
}

type nikeShirtSpy struct {
	shirtSpy
}
