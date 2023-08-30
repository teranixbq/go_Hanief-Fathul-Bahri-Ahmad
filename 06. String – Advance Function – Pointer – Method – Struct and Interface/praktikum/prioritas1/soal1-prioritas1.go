package main

import "fmt"


type Car struct {
	CarType string
	Fuelin  float64
}

func (c Car) estdistance() float64 {
	estimasifuel := 1.5 
	distance := c.Fuelin * estimasifuel
	return distance
}

func main() {
	car := Car{
		CarType: "SUV",
		Fuelin:  10.5,
	}
	jarak := car.estdistance()
	fmt.Println(car)
	fmt.Printf("car type: %s , est. distance: %.2f\n", car.CarType, jarak)
}