package bridge

type Computer interface {
	Print(text string)
	SetPrinter(Printer)
}
