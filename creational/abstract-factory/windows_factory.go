package abstractfactory

type Windows struct {
}

func (w *Windows) CreateButton() IButton {
	return &WindowsButton{
		Button: Button{Color: "Windows color", Size: "Windows size"},
	}
}

func (w *Windows) CreateMenu() IMenu {
	return &WindowsMenu{
		Menu: Menu{Header: "Windows menu", Position: "top"},
	}
}
