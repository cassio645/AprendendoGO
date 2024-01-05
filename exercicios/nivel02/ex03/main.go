/*
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
*/

package main

import "fmt"

// Constantes tipadas
const (
	a int     = 10
	b string  = "ola"
	c float64 = 2.0
)

// Constantes não tipadas
const (
	aA = 10
	bB = "ola"
	cC = 2.0
)

func main() {
	// Demonstrando os valores
	fmt.Printf("%v\t%v\t%v\n", a, b, c)
	fmt.Printf("%v\t%v\t%v\n", aA, bB, cC)
}
