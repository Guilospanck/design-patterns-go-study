package implementations

import (
	"testing"
)

func Test_Button_Press(t *testing.T) {
	t.Run("should execute command underlying function on press", func(t *testing.T) {
		// arrange
		sut := makeButtonSut()
		sut.command.On("Execute").Return()

		// act
		sut.usecase.Press()

		// assert
		sut.command.AssertNumberOfCalls(t, "Execute", 1)
		sut.command.AssertExpectations(t)
	})
}

type buttonSut struct {
	usecase *Button
	command *CommandSpy
}

func makeButtonSut() buttonSut {
	command := NewCommandSpy()

	usecase := NewButton(command)

	return buttonSut{usecase, command}
}
