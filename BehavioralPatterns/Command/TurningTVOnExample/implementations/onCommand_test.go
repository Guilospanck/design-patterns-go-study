package implementations

import "testing"

func Test_OnCommand_Execute(t *testing.T) {
	t.Run("should execute correctly underlying oncommand", func(t *testing.T) {
		// arrange
		sut := makeOnCommandSut()
		sut.device.On("TurnOn").Return()

		// act
		sut.usecase.Execute()

		// assert
		sut.device.AssertNumberOfCalls(t, "TurnOn", 1)
		sut.device.AssertExpectations(t)
	})
}

type onCommandSut struct {
	usecase *OnCommand
	device  *DeviceSpy
}

func makeOnCommandSut() onCommandSut {
	device := NewDeviceSpy()
	usecase := NewOnCommand(device)

	return onCommandSut{usecase, device}
}
