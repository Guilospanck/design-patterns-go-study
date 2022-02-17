package main

func main() {
	// component
	component := NewComponent()

	button := NewButton(component)
	checkbox := NewCheckbox(component)
	textbox := NewTextbox(component)

	// click
	button.Click()
	checkbox.Click()
	textbox.Click()

}
