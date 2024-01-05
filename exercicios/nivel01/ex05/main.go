/*
- Utilizando a solução do exercício anterior:
    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
        1. Isto já deve estar feito:
            1. Demonstre o valor da variável "x"
            2. Demonstre o tipo da variável "x"
            3. Atribua 42 à variável "x" utilizando o operador "="
            4. Demonstre o valor da variável "x"
        2. Agora faça tambem:
            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
            2. Demonstre o valor de "y"
            3. Demonstre o tipo de "y"
*/

package main

import (
	"fmt"
)

// (feito no anterior) Criando um tipo 'numeral' subjacente de int
type numeral int

// (feito no anterior) Criando uma variavel x usando var com o tipo criado
var x numeral

// Criando a variavel y com o tipo subjacente da variavel 'numeral' ou seja 'int'
var y int

func main() {
	// (feito no anterior) Demonstrando o valor de x e o tipo de x
	fmt.Printf("Valor de x = %v\n", x)
	fmt.Printf("Tipo de x = %T\n", x)

	// (feito no anterior) Atribuindo 42 a x e demonstrando o valor
	x = 42
	fmt.Printf("Valor de x = %v\n", x)

	// Convertendo x para int e atribuindo o valor de x para y
	y = int(x)

	// Demonstrando o valor de y e o tipo de y
	fmt.Printf("Valor de y = %v\n", y)
	fmt.Printf("Tipo de y = %T\n", y)

}
