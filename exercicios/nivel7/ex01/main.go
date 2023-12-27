/*
- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.
*/

package main

import "fmt"

func main() {
	valor := 10

	// Demonstrando o valor e o endereço da memoria do valor
	fmt.Printf("Valor: %d, está no endereço: %d\n", valor, &valor)
}
