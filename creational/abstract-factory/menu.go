package abstractfactory

import "fmt"

type IMenu interface {
	Show()
	Hide()
}

type Menu struct {
	Header   string
	Position string
}

func (m *Menu) Show() {
	fmt.Printf("Menu apper with header %s appears in %s position\n", m.Header, m.Position)
}

func (m *Menu) Hide() {
	fmt.Printf("Hide menu with header %s\n", m.Header)
}
