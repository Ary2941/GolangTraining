package main

import "fmt"

func rocambole(texto string) (resultado string) {

	resultado = texto + "é um texto"
	return
}

func main() {

	qre := rocambole("ato da compadrecida")
	fmt.Println(qre)
}
