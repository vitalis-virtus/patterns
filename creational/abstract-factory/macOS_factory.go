package abstractfactory

type MacOS struct {
}

func (m *MacOS) CreateButton() IButton {
	return &MacOSButton{Button: Button{Color: "MacOS color", Size: "MacOS size"}}
}

func (m *MacOS) CreateMenu() IMenu {
	return &MacOSMenu{Menu: Menu{Header: "MacOS menu", Position: "left"}}
}
