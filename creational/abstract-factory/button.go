package abstractfactory

import "fmt"

type IButton interface {
	PrintColor()
	PrintSize()
}

type Button struct {
	Color string
	Size  string
}

func (b *Button) PrintColor() {
	fmt.Printf("Button color: %s\n", b.Color)
}

func (b *Button) PrintSize() {
	fmt.Printf("Button size: %s\n", b.Size)
}
