package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

func (m *Mac) Print(text string) {
	fmt.Println("Print request for mac")
	m.printer.Print(text)
}
