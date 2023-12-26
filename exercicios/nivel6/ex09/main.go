// Callback: passe uma função como argumento a outra função.

package main

import "fmt"

func main() {

	// Chamando a funcao que recebe como parâmetro outra função
	totalImpares := somarImpares(somarTudo, []int{25, 32, 11, 9, 26, 31, 94, 10, 65}...)
	fmt.Println(totalImpares)

}

// Função que será usada no callback
func somarTudo(slice ...int) int {
	total := 0
	for _, valor := range slice {
		total += valor
	}
	return total
}

// Função que separa os valores impares e chama outra função para somar os valores
func somarImpares(f func(x ...int) int, slice ...int) int {
	var sliceImpar []int

	for _, numero := range slice {
		if numero%2 == 1 {
			sliceImpar = append(sliceImpar, numero)
		}
	}

	valorTotal := f(sliceImpar...)
	return valorTotal
}
