/*
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.
*/

package main

import "fmt"

func main() {

	// Criando um array vazio do tipo int
	var myArray = [5]int{}

	// Atribuindo valores para os indices
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50

	// Demonstrando os valores com range
	for index, valor := range myArray {
		fmt.Printf("Posição %d: Valor: %d\n", index, valor)
	}

	// Demonstrando o tipo do array
	fmt.Printf("Tipo do Array: %T\n", myArray)

}
