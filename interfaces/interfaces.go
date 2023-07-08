package main

import "fmt"

type MotorVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance float64
	avgspeed string
	fuel     float64
}

type Audi struct {
	distance float64
	fuel     float64
}

func (b BMW) AverageSpeed() string {
	return b.avgspeed
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func totalMileage(m MotorVehicle) {
	au := m.(BMW)
	fmt.Printf(au.AverageSpeed())
}

func main() {
	b1 := BMW{
		distance: 167.9,
		fuel:     36,
		avgspeed: "58",
	}
	totalMileage(b1)
}
