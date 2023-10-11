package main

import (
	"fmt"
	"log"
	abstractfactory "patterns/creational/abstract-factory"
	"patterns/creational/abstract-factory/sports"
	"patterns/creational/builder"
	"patterns/creational/factory"
	"runtime"
)

func main() {
	// fabric()

	// abstract_fabric()

	// abstract_fabric_sports()

	builder_pattern()
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

func abstract_fabric_sports() {
	// nike factory
	nikeFactory, err := sports.GetSports("nike")
	if err != nil {
		fmt.Println(err)
	}

	nikeShirt := nikeFactory.CreateShirt()
	fmt.Println(nikeShirt.GetColor())

	nikeShirt.SetColor("blue")
	fmt.Println(nikeShirt.GetColor())

	nikeShoe := nikeFactory.CreateShoe()
	fmt.Println(nikeShoe.GetLogo())

	nikeShoe.SetLogo("nike black")
	fmt.Println(nikeShoe.GetLogo())

	// adidas factory
	adidasFactory, err := sports.GetSports("adidas")
	if err != nil {
		fmt.Println(err)
	}
	adidasShirt := adidasFactory.CreateShirt()
	fmt.Println(adidasShirt.GetColor())

	adidasShirt.SetColor("white")
	fmt.Println(adidasShirt.GetColor())

	adidasShoe := adidasFactory.CreateShoe()
	fmt.Println(adidasShoe.GetLogo())

	adidasShoe.SetLogo("adidas retro")
	fmt.Println(adidasShoe.GetLogo())

	// acics
	acicsFactory, err := sports.GetSports("acics")
	if err != nil {
		log.Fatal(err)
	}

	acicsShoe := acicsFactory.CreateShoe()
	fmt.Println(acicsShoe.GetLogo())
}

func builder_pattern() {
	carBuilder := builder.NewCarBuilder()

	car := carBuilder.SetColor("green").SetEngineType("V8").SetNavigation(false).SetSunroof(true)

	fmt.Println(car.Build())
}
