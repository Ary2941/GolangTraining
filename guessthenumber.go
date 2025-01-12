package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var userInput int

	var houdini int = rand.Intn(10-1) + 1

	fmt.Println("Guess the number! you have 3 tries!!")

	var counter = 0
	for counter < 3 {
		fmt.Scan(&userInput)
		if userInput == houdini {
			fmt.Println("You found out!")
			return
		} else {
			fmt.Println("Nah, try again!")
		}
		counter++
	}
	fmt.Printf("Nice tries buddy, the right number was %v", houdini)
}
