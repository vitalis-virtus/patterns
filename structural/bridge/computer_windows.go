package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (m *Windows) SetPrinter(printer Printer) {
	m.printer = printer
}

func (m *Windows) Print(text string) {
	fmt.Println("Print request for windows")
	m.printer.Print(text)
}
