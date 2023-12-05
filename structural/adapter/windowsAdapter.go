package adapter

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func NewWindowsAdapter(w *Windows) *WindowsAdapter {
	return &WindowsAdapter{w}
}

func (w *WindowsAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lightning signal to USB")
	w.windowsMachine.InsertIntoUSBPort()
}
