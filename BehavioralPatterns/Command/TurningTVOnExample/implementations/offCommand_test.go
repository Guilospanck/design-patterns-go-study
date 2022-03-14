package implementations

import "testing"

func Test_OffCommand_Execute(t *testing.T) {
	t.Run("should execute correctly function under offcommand", func(t *testing.T) {
		// arrange
		sut := makeOffCommandSut()
		sut.device.On("TurnOff").Return()

		// act
		sut.usecase.Execute()

		// assert
		sut.device.AssertNumberOfCalls(t, "TurnOff", 1)
		sut.device.AssertExpectations(t)
	})
}

type offCommandSut struct {
	usecase *OffCommand
	device  *DeviceSpy
}

func makeOffCommandSut() offCommandSut {
	device := NewDeviceSpy()
	usecase := NewOffCommand(device)

	return offCommandSut{usecase, device}
}
