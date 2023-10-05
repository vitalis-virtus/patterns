package abstractfactory

type IGUI interface {
	CreateButton() IButton
	CreateMenu() IMenu
}

func GetGUI(system string) IGUI {
	switch system {
	case "windows":
		return &Windows{}
	case "macOS":
		return &MacOS{}
	}
	return nil
}
