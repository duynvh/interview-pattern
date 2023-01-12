package main

import "fmt"

func main() {
	manufacturingVehicle := ManufacturingDirector{}
	bicycleBuilder := &BicycleBuilder{}

	manufacturingVehicle.SetBuilder(bicycleBuilder)
	manufacturingVehicle.Construct()

	bicycle := bicycleBuilder.GetVehicle()
	fmt.Printf("Vehicle is %s has %d wheels, %d seats.", bicycle.Structure, bicycle.Wheels, bicycle.Seats)
}
