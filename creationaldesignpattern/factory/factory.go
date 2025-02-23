package main

import "fmt"

type Car interface {
	getCarType() string
}

type SedanCar struct {
	Name string
}

func getSedanCar() *SedanCar {
	return &SedanCar{
		Name: "Honda City",
	}
}

func (sc *SedanCar) getCarType() string {
	return sc.Name
}

type HatchBackCar struct {
	Name string
}

func getHatchBackCar() *HatchBackCar {
	return &HatchBackCar{
		Name: "Polo",
	}
}

func (hc *HatchBackCar) getCarType() string {
	return hc.Name
}

func main() {
	car1 := getCarFactory(1)
	fmt.Println("Car 1 type: ", car1.getCarType())
	car2 := getCarFactory(2)
	fmt.Println("Car 2 type: ", car2.getCarType())
}

func getCarFactory(cType int) Car {
	var car Car
	switch cType {
	case 1:
		car = getSedanCar()
	case 2:
		car = getHatchBackCar()
	}

	return car
}
