/*
- Utilizando este código: https://play.golang.org/p/MvL6uamrJP
- ...use um select statement para demonstrar os valores do canal.
*/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

// Alterado de func gen(q <- chan int) Para func gen(q chan<- int)
// Para q poder receber valores
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // encerrando o canal c quando todos valores forem recebidos
		q <- 0   // atribuindo um valor ao canal q
	}()

	return c
}

// Função que mostra os valores contidos em C e sai quando q recebe algum valor
func receive(c, q <-chan int) {
	for {
		select {
		case valor := <-c:
			fmt.Println(valor)
		case <-q:
			return
		}
	}
}
