package main

import "fmt"

type Vehicle struct {
	Brand string
}

func (v Vehicle) Move() {
	fmt.Printf("%s moving generically\n", v.Brand)
}

type Car struct {
	Vehicle // embedding: expõe métodos/campos de Vehicle
	Doors   int
}

func (c Car) Move() {
	fmt.Printf("%s car driving with %d doors\n", c.Brand, c.Doors)
}

type Bicycle struct {
	Vehicle
}

func main() {
	vehicle := Vehicle{Brand: "Generic"}
	vehicle.Move()

	car := Car{
		Vehicle: Vehicle{Brand: "Tesla"},
		Doors:   4,
	}
	car.Move()

	bike := Bicycle{Vehicle: Vehicle{Brand: "Caloi"}}
	bike.Move() // usa implementação de Vehicle
	bike.Brand = "Sense"
	bike.Move() // campo promovido pelo embedding
}
