/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func main() {
	// Passando varios inteiros
	somado := somarTudo(2, 5, 3, 2, 10, 11, 14, 45, 20)
	fmt.Println("Todos os inteiros somaram:", somado)

	// Passando um slice
	slice := []int{2, 10, 54, 32, 12, 44, 12}
	somadoSlice := somarTudo(slice...) // função que recebe int
	fmt.Println("Todos os valores do slice somaram:", somadoSlice)

	somadoSlice_2 := somarSlice(slice) // funcao que recebe slice
	fmt.Println("Todos os valores do slice somaram:", somadoSlice_2)
}

// Função que recebe um parâmetro variádico e retorna a soma de todos os ints recebidos
func somarTudo(inteiros ...int) int {
	total := 0
	for _, valor := range inteiros {
		total += valor
	}
	return total
}

// Função que recebe um SLICE OF INT e soma todos os valores
func somarSlice(inteiros []int) int {
	total := 0
	for _, valor := range inteiros {
		total += valor
	}
	return total
}
