package bridge

import "fmt"

type HP struct {
}

func (p *HP) Print(text string) {
	fmt.Println("Printing by a HP Printer")
	fmt.Println(text)
}
