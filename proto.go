package main

import (
	"fmt"
)

type electricEngine struct {
	milesPerKWh float64
	batteryKwh  float64
}

type gasEngine struct {
	milesPerGallon float64
	gasLeft        float64
}

func (e electricEngine) milesLeft() float64 {
	return e.milesPerKWh * e.batteryKwh
}
func (e gasEngine) milesLeft() float64 {
	return e.milesPerGallon * e.gasLeft
}

// writing generic function with interfaces

type engine interface {
	milesLeft() float64
}

func canMakeIt(e engine, milesTo float64) bool {
	if e.milesLeft() > milesTo {
		return true
	}
	return false
}

var tesla = electricEngine{milesPerKWh: 100, batteryKwh: 12}
var toyota = gasEngine{milesPerGallon: 200, gasLeft: 20}

func main() {
	fmt.Println("canMakeIt(tesla, 200): ", canMakeIt(tesla, 200))
	fmt.Println("canMakeIt(toyota, 200)", canMakeIt(toyota, 200))
}
