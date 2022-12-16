package main

import "fmt"

func main() {
	GetInstance().Increase()
	GetInstance().Increase()

	fmt.Println(GetInstance().Get())
}
