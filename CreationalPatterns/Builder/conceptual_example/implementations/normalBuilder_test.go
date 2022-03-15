package implementations

import (
	"base/CreationalPatterns/Builder/conceptual_example/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NormalBuilder_SetWindowType(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeSut()

		// act
		sut.normalBuilder.SetWindowType()

		// assert
		assert.Equal(t, sut.windowType, sut.normalBuilder.windowType)
	})
}

func Test_NormalBuilder_SetDoorType(t *testing.T) {
	t.Run("should set door type", func(t *testing.T) {
		// arrange
		sut := makeSut()

		// act
		sut.normalBuilder.SetDoorType()

		// assert
		assert.Equal(t, sut.doorType, sut.normalBuilder.doorType)
	})
}

func Test_NormalBuilder_SetNumFloor(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeSut()

		// act
		sut.normalBuilder.SetNumFloor()

		// assert
		assert.Equal(t, sut.floor, sut.normalBuilder.floor)
	})
}

func Test_NormalBuilder_GetHouse(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeSut()
		sut.normalBuilder.doorType = sut.doorType
		sut.normalBuilder.windowType = sut.windowType
		sut.normalBuilder.floor = sut.floor

		// act
		result := sut.normalBuilder.GetHouse()

		// assert
		assert.Equal(t, sut.house, result)
	})
}

type normalBuilderSut struct {
	normalBuilder        *normalBuilder
	windowType, doorType string
	floor                int
	house                domain.House
}

func makeSut() normalBuilderSut {
	normalBuilder := NewNormalBuilder()
	windowType := "Wooden Window"
	doorType := "Wooden Door"
	floor := 2

	house := domain.House{
		WindowType: windowType,
		DoorType:   doorType,
		Floor:      floor,
	}

	return normalBuilderSut{normalBuilder, windowType, doorType, floor, house}
}
