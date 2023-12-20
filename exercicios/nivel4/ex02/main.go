/*
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
*/

package main

import "fmt"

func main() {

		// Criando um slice vazio
		mySlice := []int{}

		// Atribuindo valores
		mySlice = append(mySlice, 10)
		mySlice = append(mySlice, 20)
		mySlice = append(mySlice, 30)
		mySlice = append(mySlice, 40)
		mySlice = append(mySlice, 50)
		mySlice = append(mySlice, 60)
		mySlice = append(mySlice, 70)
		mySlice = append(mySlice, 80)
		mySlice = append(mySlice, 90)
		mySlice = append(mySlice, 100)
	
		// Demonstrando os valores com range
		for index, valor := range mySlice {
			fmt.Printf("Posição %d: Valor: %d\n", index, valor)
		}
	
		// Demonstrando o tipo do slice
		fmt.Printf("Tipo do Slice: %T\n", mySlice)
	
}