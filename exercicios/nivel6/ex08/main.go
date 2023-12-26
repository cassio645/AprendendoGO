/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/

package main

import "fmt"

func main() {
	// x recebe a chamada da funcao
	x := chamada()

	x() // chamando a outra funcao

}

// Função que retorna uma função
func chamada() func() {
	return func() {
		fmt.Println("Olá mundo!")
	}
}