/*
- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
*/

package main

import "fmt"

// Criando um tipo 'numeral' subjacente de int
type numeral int

// Criando uma variavel x usando var com o tipo criado
var x numeral

func main() {
	// Demonstrando o valor de x e o tipo de x
	fmt.Printf("Valor de x = %v\n", x)
	fmt.Printf("Tipo de x = %T\n", x)

	// Atribuindo 42 a x e demonstrando o valor
	x = 42
	fmt.Printf("Valor de x = %v\n", x)
}
