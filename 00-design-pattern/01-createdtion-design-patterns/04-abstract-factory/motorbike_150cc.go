package main

type Motorbike150CC struct{}

func (*Motorbike150CC) GetCC() string {
	return "155cc"
}

func (*Motorbike150CC) NumWheels() int {
	return 2
}

func (*Motorbike150CC) NumSeats() int {
	return 2
}
