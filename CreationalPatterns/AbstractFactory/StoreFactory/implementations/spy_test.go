package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Spy_ShoeSpy_SetLogo(t *testing.T) {
	t.Run("Should call SetLogo with logo", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shoeSpy.On("SetLogo", sut.logo).Return()

		// act
		sut.shoeSpy.SetLogo(sut.logo)

		// assert
		sut.shoeSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShoeSpy_SetSize(t *testing.T) {
	t.Run("Should call SetSize with size", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shoeSpy.On("SetSize", sut.size).Return()

		// act
		sut.shoeSpy.SetSize(sut.size)

		// assert
		sut.shoeSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShoeSpy_GetSize(t *testing.T) {
	t.Run("Should call GetSize with size", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shoeSpy.On("GetSize").Return(sut.size)

		// act
		result := sut.shoeSpy.GetSize()

		// assert
		assert.NotNil(t, result)
		assert.Equal(t, sut.size, result)
		sut.shoeSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShoeSpy_GetLogo(t *testing.T) {
	t.Run("Should call GetLogo with logo", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shoeSpy.On("GetLogo").Return(sut.logo)

		// act
		result := sut.shoeSpy.GetLogo()

		// assert
		assert.NotNil(t, result)
		assert.Equal(t, sut.logo, result)
		sut.shoeSpy.AssertExpectations(t)
	})
}

//

func Test_Spy_ShirtSpy_SetLogo(t *testing.T) {
	t.Run("Should call SetLogo with logo", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shirtSpy.On("SetLogo", sut.logo).Return()

		// act
		sut.shirtSpy.SetLogo(sut.logo)

		// assert
		sut.shirtSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShirtSpy_SetSize(t *testing.T) {
	t.Run("Should call SetSize with size", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shirtSpy.On("SetSize", sut.size).Return()

		// act
		sut.shirtSpy.SetSize(sut.size)

		// assert
		sut.shirtSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShirtSpy_GetSize(t *testing.T) {
	t.Run("Should call GetSize with size", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shirtSpy.On("GetSize").Return(sut.size)

		// act
		result := sut.shirtSpy.GetSize()

		// assert
		assert.NotNil(t, result)
		assert.Equal(t, sut.size, result)
		sut.shirtSpy.AssertExpectations(t)
	})
}

func Test_Spy_ShirtSpy_GetLogo(t *testing.T) {
	t.Run("Should call GetLogo with logo", func(t *testing.T) {
		// arrange
		sut := makeSpySut()
		sut.shirtSpy.On("GetLogo").Return(sut.logo)

		// act
		result := sut.shirtSpy.GetLogo()

		// assert
		assert.NotNil(t, result)
		assert.Equal(t, sut.logo, result)
		sut.shirtSpy.AssertExpectations(t)
	})
}

//

type spySut struct {
	shoeSpy  *shoeSpy
	shirtSpy *shirtSpy
	logo     string
	size     int
}

func makeSpySut() spySut {
	logo := "testLogo"
	size := 20

	shoeSpy := NewShoeSpy(size, logo)
	shirtSpy := NewShirtSpy(size, logo)

	return spySut{shoeSpy, shirtSpy, logo, size}
}
