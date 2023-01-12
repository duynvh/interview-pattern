package main

import "fmt"

const (
	latteName      = "Latte"
	cappuchinoName = "Cappuchino"
)

func GetCoffeeDrink(name string) (ICoffeeDrink, error) {
	switch name {
	case latteName:
		return newLatte(), nil
	case cappuchinoName:
		return newCappuchino(), nil
	default:
		return nil, fmt.Errorf("Unknown coffee drink: %s", name)
	}
}
