package main

type CoffeeDrink struct {
    name string
}

type ICoffeeDrink interface {
    GetName() string
}

func (me CoffeeDrink) GetName() string {
    return me.name
}