package main

import "fmt"

func main() {
	/*
		Example Abstract Factory
	*/
	fmt.Println("*** Example Abstract Factory ***")

	bicycleFactory, err := BuildFactory(BicycleFactoryType)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	sportBicycle, err := bicycleFactory.NewVehicle(SportBicycleType)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Sport Bicycle:")
	fmt.Printf("Vehicle has %d wheels, %d seats.\n", sportBicycle.NumWheels(), sportBicycle.NumSeats())

	fmt.Print("*** End of Abstract Factory ***\n\n\n")
}
