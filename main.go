package main

import (
	"fmt"
	"log"
	abstractfactory "patterns/creational/abstract-factory"
	"patterns/creational/abstract-factory/sports"
	"patterns/creational/builder"
	"patterns/creational/factory"
	"patterns/creational/prototype"
	"patterns/creational/singletone"
	"patterns/structural/adapter"
	"patterns/structural/bridge"
	"runtime"
)

func main() {
	// fabric()

	// abstract_fabric()

	// abstract_fabric_sports()

	// builder_pattern()

	// prototype_pattern()

	// singletone_pattern()

	// adapter_pattern()

	brindge_pattern()
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
	carBuilderSabre := builder.NewCarBuilder()

	car := carBuilderSabre.SetColor("green").SetEngineType("V8").SetNavigation(false).SetSunroof(true)
	sabre := car.Build()

	director := builder.NewDirector(builder.NewCarBuilder())
	camaro := director.ConstructCar("red", "V12", false, true)

	fmt.Println(sabre)
	fmt.Println(camaro)
}

func prototype_pattern() {
	file1 := prototype.File{Name: "File1"}
	file2 := prototype.File{Name: "File2"}
	file3 := prototype.File{Name: "File3"}

	folder1 := prototype.Folder{
		Name:     "Folder1",
		Children: []prototype.Node{&file1},
	}

	folder2 := prototype.Folder{
		Name:     "Folder2",
		Children: []prototype.Node{&folder1, &file2, &file3},
	}

	folder2.Print()

	cloneFolder := folder2.Clone()
	cloneFolder.Print()
}

func singletone_pattern() {
	for i := 0; i < 10; i++ {
		singletone.GetInstance()
	}
}

// structural patterns

func adapter_pattern() {
	// creating new client
	client := adapter.Client{}

	// creating new Mac machinec
	macMachine := adapter.Mac{}
	client.InsertLigtingConnectorIntoComputer(&macMachine)

	// creating new Windows machines
	windowsMachine := adapter.Windows{}
	// creating new WindowsAdapter
	windowsAdapter := adapter.NewWindowsAdapter(&windowsMachine)
	client.InsertLigtingConnectorIntoComputer(windowsAdapter)
}

func brindge_pattern() {
	hpPrinter := bridge.HP{}
	epsonPrinter := bridge.Epson{}

	macComputer := bridge.Mac{}
	windowsComputer := bridge.Windows{}

	macComputer.SetPrinter(&hpPrinter)
	macComputer.Print("hello world")
	macComputer.SetPrinter(&epsonPrinter)
	macComputer.Print("hello world")

	windowsComputer.SetPrinter(&hpPrinter)
	windowsComputer.Print("hello world")
	windowsComputer.SetPrinter(&epsonPrinter)
	windowsComputer.Print("hello world")
}
