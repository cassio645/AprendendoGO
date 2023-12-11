/*
- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
    1. Identificador "x" deverá ter tipo int
    2. Identificador "y" deverá ter tipo string
    3. Identificador "z" deverá ter tipo bool
- Na função main:
    1. Demonstre os valores de cada identificador
    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
*/

package main

import "fmt"

// Declarando as variáveis sem atribuir valores
var x int
var y string
var z bool

func main() {
	// Demonstrando os valores de cada variável
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// O compilador atribuiu o valor 'zero' referente a cada tipo.
	// valor 'zero' do int é 0.
	// valor 'zero' da string é 	.
	// valor 'zero' do bool é false.
}
