package factory

import "fmt"

type Transport struct {
	name   string
	wheels int
}

func (t *Transport) Ride() {
	fmt.Printf("%s riding on the %v wheels\n", t.name, t.wheels)
}
