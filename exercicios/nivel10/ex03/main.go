/*
- Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
- ...use um for range loop para demonstrar os valores do canal.
*/

package main

import (
	"fmt"
)

func main() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	// Colocando o for dentro de uma go func auto-executavel para poder executar paralelamente
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		// Fechando o canal após adicionar todos os elementos
		close(c)
	}()

	return c
}

// Função receive criada para demonstrar os valores do canal
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
