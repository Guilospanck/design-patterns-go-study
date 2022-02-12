package main

func main() {
	// Create our receivers (devices)
	tv := NewTV()
	radio := NewRadio()

	// Create our commands
	onTVCommand := NewOnCommand(tv)
	onRadioCommand := NewOnCommand(radio)
	offTVCommand := NewOffCommand(tv)
	offRadioCommand := NewOffCommand(radio)

	// Create our sender (Button)
	onTVButton := NewButton(onTVCommand)
	offTVButton := NewButton(offTVCommand)
	onRadioButton := NewButton(onRadioCommand)
	offRadioButton := NewButton(offRadioCommand)

	// press buttons
	onTVButton.Press()
	offTVButton.Press()
	onRadioButton.Press()
	offRadioButton.Press()

}
