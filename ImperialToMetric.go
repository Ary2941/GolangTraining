package main

import (
	"fmt"
)

const inchProportion float64 = 0.393701

var UserInput float64

func main() {
	fmt.Print("digite o valor em centímetros: ")
	fmt.Scan(&UserInput)

	var valueInInch float64 = UserInput * inchProportion
	var valueInFeet float64 = (UserInput * inchProportion) / 12

	fmt.Printf("%.2f inches or %.2f feet", valueInInch, valueInFeet)
}
