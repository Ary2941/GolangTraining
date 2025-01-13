package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var phrase string

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Remove any surrounding whitespace including the newline.
	fmt.Printf("You entered: %s", text)
}
