package main

import (
	"log"
	abstractfactory "patterns/creational/abstract-factory"
	"patterns/creational/factory"
	"runtime"
)

func main() {
	// fabric()

	// abstract_fabric()
}

// creational patterns

func fabric() {
	car, err := factory.GetTransport("car")
	if err != nil {
		log.Fatal(err)
	}
	car.Ride()

	bicycle, err := factory.GetTransport("bicycle")
	if err != nil {
		log.Fatal(err)
	}
	bicycle.Ride()

	scooter, err := factory.GetTransport("scooter")
	if err != nil {
		log.Fatal(err)
	}
	scooter.Ride()
}

func abstract_fabric() {
	factory := abstractfactory.GetGUI(runtime.GOOS)

	button := factory.CreateButton()
	button.PrintColor()
	button.PrintSize()

	menu := factory.CreateMenu()
	menu.Show()
	menu.Hide()
}
