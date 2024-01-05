/*
 	- Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
*/

package main

import "fmt"

func main() {

	// Atribuindo um valor int e demontrando em decimal, binario e hexadecimal
	valor := 100
	fmt.Printf("Decimal: %d\tBinário: %b\tHexadecimal: %#x\n", valor, valor, valor)

	// Deslocando 1 bit para esquerda e demonstrando o novo valor
	valorDeslocado := valor << 1
	fmt.Printf("Decimal: %d\tBinário: %b\tHexadecimal: %#x\n", valorDeslocado, valorDeslocado, valorDeslocado)

}
