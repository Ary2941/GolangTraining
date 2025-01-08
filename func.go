package main

import (
	"fmt"
)

func functionNoArgs() {
	fmt.Println("this function has no args")
}

func functionOneArg(arg string) {
	fmt.Printf("this function has an argument %v\n", arg)
}

func functiontwoArg(arg string, age int) {
	fmt.Printf("this function has two arguments: %v", arg)
	fmt.Printf(" and %v", age)
}

func main() {
	functionNoArgs()
	functionOneArg("rocamboole")
	functiontwoArg("rocambole", 94)
}
