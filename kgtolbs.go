package main

import (
	"fmt"
)

var kg = 65.0

const lbs = 2.20462262185
const oz = 35.274

func main() {
	fmt.Printf("%v kg is equivalent to ", kg)
	fmt.Printf("%v lbs", kg*lbs)
	fmt.Print(" or ")
	fmt.Printf("%v oz", kg*oz)
}
