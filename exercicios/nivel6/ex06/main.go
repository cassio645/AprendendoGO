// - Crie e utilize uma função anônima.
package main

import "fmt"

func main() {

	// Função anonima que demonstra o dobro do valor recebido
	func(x int) {
		fmt.Println("O dobro de", x, "é", x*2)
	}(15)

}
