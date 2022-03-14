package implementations

import "testing"

func Test_CommmandSpy_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		// arrange
		sut := NewCommandSpy()
		sut.On("Execute").Return()

		// act
		sut.Execute()

		// assert
		sut.AssertExpectations(t)
	})
}

func Test_DeviceSpy(t *testing.T) {
	t.Run("should execute correctly turn on method", func(t *testing.T) {
		// arrange
		sut := NewDeviceSpy()
		sut.On("TurnOn").Return()

		// act
		sut.TurnOn()

		// assert
		sut.AssertExpectations(t)
	})

	t.Run("should execute correctly turn off method", func(t *testing.T) {
		// arrange
		sut := NewDeviceSpy()
		sut.On("TurnOff").Return()

		// act
		sut.TurnOff()

		// assert
		sut.AssertExpectations(t)
	})
}
