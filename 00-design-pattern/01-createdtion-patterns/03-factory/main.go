package main

import "fmt"

func main() {
	cappuchino, err := GetCoffeeDrink("Cappuchino")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cappuchino.GetName())
	}

	latte, err := GetCoffeeDrink("Latte")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(latte.GetName())
	}

	_, err = GetCoffeeDrink("Error")
	if err != nil {
		fmt.Println(err)
	}
}
