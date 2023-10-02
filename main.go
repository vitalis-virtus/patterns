package main

import (
	"log"
	"patterns/creational/factory"
)

func main() {
	fabric()
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
