// - Faça esse código funcionar: https://play.golang.org/p/oB-p3KMiH6

package main

import (
	"fmt"
)

func main() {

	// Fazendo o canal ser bidirecional, que pode tanto receber quanto enviar valores
	cs := make(chan int)
	// cs := make(chan<- int)   - antigo, canal apenas de envio, não aceita receber valores

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
