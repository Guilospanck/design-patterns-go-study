package main

import "base/BehavioralPatterns/Command/TurningTVOnExample/implementations"

func main() {
	// Create our receivers (devices)
	tv := implementations.NewTV()
	radio := implementations.NewRadio()

	// Create our commands
	onTVCommand := implementations.NewOnCommand(tv)
	onRadioCommand := implementations.NewOnCommand(radio)
	offTVCommand := implementations.NewOffCommand(tv)
	offRadioCommand := implementations.NewOffCommand(radio)

	// Create our sender (Button)
	onTVButton := implementations.NewButton(onTVCommand)
	offTVButton := implementations.NewButton(offTVCommand)
	onRadioButton := implementations.NewButton(onRadioCommand)
	offRadioButton := implementations.NewButton(offRadioCommand)

	// press buttons
	onTVButton.Press()
	offTVButton.Press()
	onRadioButton.Press()
	offRadioButton.Press()

}
