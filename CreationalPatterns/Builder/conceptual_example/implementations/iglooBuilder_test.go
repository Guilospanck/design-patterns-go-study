package implementations

import (
	"base/CreationalPatterns/Builder/conceptual_example/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IglooBuilder_SetWindowType(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeIglooBuilderSut()

		// act
		sut.iglooBuilder.SetWindowType()

		// assert
		assert.Equal(t, sut.windowType, sut.iglooBuilder.windowType)
	})
}

func Test_IglooBuilder_SetDoorType(t *testing.T) {
	t.Run("should set door type", func(t *testing.T) {
		// arrange
		sut := makeIglooBuilderSut()

		// act
		sut.iglooBuilder.SetDoorType()

		// assert
		assert.Equal(t, sut.doorType, sut.iglooBuilder.doorType)
	})
}

func Test_IglooBuilder_SetNumFloor(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeIglooBuilderSut()

		// act
		sut.iglooBuilder.SetNumFloor()

		// assert
		assert.Equal(t, sut.floor, sut.iglooBuilder.floor)
	})
}

func Test_IglooBuilder_GetHouse(t *testing.T) {
	t.Run("should set window type", func(t *testing.T) {
		// arrange
		sut := makeIglooBuilderSut()
		sut.iglooBuilder.doorType = sut.doorType
		sut.iglooBuilder.windowType = sut.windowType
		sut.iglooBuilder.floor = sut.floor

		// act
		result := sut.iglooBuilder.GetHouse()

		// assert
		assert.Equal(t, sut.house, result)
	})
}

type iglooBuilderSut struct {
	iglooBuilder         *iglooBuilder
	windowType, doorType string
	floor                int
	house                domain.House
}

func makeIglooBuilderSut() iglooBuilderSut {
	iglooBuilder := NewIglooBuilder()
	windowType := "Snow Window"
	doorType := "Snow Door"
	floor := 1

	house := domain.House{
		WindowType: windowType,
		DoorType:   doorType,
		Floor:      floor,
	}

	return iglooBuilderSut{iglooBuilder, windowType, doorType, floor, house}
}
