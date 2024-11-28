package main

import "fmt"

// Application that contains the actual logic
type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (a *SedanType) getCar() string {
	return "Honda City"
}

type HatchbackType struct {
	Name string
}

func getNewHatchback() *HatchbackType {
	return &HatchbackType{}
}

func (h *HatchbackType) getCar() string {
	return "Golf GTI"
}

// Client
func main() {
	getCarFactory(1)
	getCarFactory(2)
}

// Factory
func getCarFactory(carType int) {
	var car Car
	if carType == 1 {
		car = getNewSedan()
	}
	if carType == 2 {
		car = getNewHatchback()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)
}
