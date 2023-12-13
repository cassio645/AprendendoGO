/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/

package main

import "fmt"

// Criando valores dos proximos anos com iota
const (
	proximoAno = iota + 2024
	doisAnos
	treAnos
	quatroAnos
)

func main() {

	// Demonstrando os valores
	fmt.Printf("Próximo ano: %v\nDaqui dois anos: %v\nDaqui três anos: %v\nDaqui quatro anos: %v\n", proximoAno, doisAnos, treAnos, quatroAnos)

}
