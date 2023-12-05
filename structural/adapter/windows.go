package adapter

import "fmt"

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into Windows machine")
}
