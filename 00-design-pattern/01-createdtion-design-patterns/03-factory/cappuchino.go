package main

type Cappuchino struct {
	CoffeeDrink
}

func newCappuchino() *Cappuchino {
	return &Cappuchino{
		CoffeeDrink: CoffeeDrink{
			name: "Cappuchino",
		},
	}
}
