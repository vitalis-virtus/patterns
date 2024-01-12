package bridge

import "fmt"

type Epson struct {
}

func (p *Epson) Print(text string) {
	fmt.Println("Printing by a Epson Printer")
	fmt.Println(text)
}
