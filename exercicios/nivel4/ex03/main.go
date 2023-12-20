/*
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/

package main

import "fmt"

func main() {

	// Criando um slice com 10 valores int
	mySlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// DEMONSTRANDO
	// Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	fmt.Println("Do primeiro ao terceiro item:", mySlice[:3])

	// Do quinto ao último item do slice (incluindo o último item!)
	fmt.Println("Do quinto ao último item:", mySlice[4:])

	// Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	fmt.Println("Do segundo ao sétimo item:", mySlice[1:7])

	// Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	fmt.Println("Do terceiro ao penúltimo item:", mySlice[2:9])

	// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	fmt.Println("O mesmo resultado usando len():", mySlice[2:len(mySlice)-1])
}
