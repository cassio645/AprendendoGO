/*
- Atribua uma função a uma variável.
- Chame essa função.
*/

package main

import "fmt"

func main() {

	// variável bomdia recebendo uma função
	bomdia := func() {
		fmt.Println("Olá bom dia!")
	}

	// chamando a função
	bomdia()

}
