package implementations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Director_SetBuilder(t *testing.T) {
	t.Run("should set builder properly", func(t *testing.T) {
		// arrange
		sut := makeDirectorSut()

		// act and assert
		assert.Equal(t, sut.normalBuilder, sut.director.builder)

		sut.director.SetBuilder(sut.iglooBuilder)
		assert.Equal(t, sut.iglooBuilder, sut.director.builder)
		sut.director.SetBuilder(sut.normalBuilder)
		assert.Equal(t, sut.normalBuilder, sut.director.builder)
	})
}

func Test_Director_BuildHouse(t *testing.T) {
	t.Run("it should build normal house properly", func(t *testing.T) {
		// arrange
		sut := makeDirectorSut()

		// act
		result := sut.director.BuildHouse()

		// assert
		assert.Equal(t, sut.normalBuilder.GetHouse(), result)
	})

	t.Run("it should build igloo house properly", func(t *testing.T) {
		// arrange
		sut := makeDirectorSut()

		// act
		sut.director.SetBuilder(sut.iglooBuilder)
		result := sut.director.BuildHouse()

		// assert
		assert.Equal(t, sut.iglooBuilder.GetHouse(), result)
	})
}

type directorSut struct {
	director      *director
	normalBuilder *normalBuilder
	iglooBuilder  *iglooBuilder
}

func makeDirectorSut() directorSut {
	normalBuilder := NewNormalBuilder()
	iglooBuilder := NewIglooBuilder()

	normalBuilder.SetWindowType()
	normalBuilder.SetDoorType()
	normalBuilder.SetNumFloor()

	iglooBuilder.SetWindowType()
	iglooBuilder.SetDoorType()
	iglooBuilder.SetNumFloor()

	director := NewDirector(normalBuilder)

	return directorSut{director, normalBuilder, iglooBuilder}
}
