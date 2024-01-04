/*
- Nível 10?! Êita! Parabéns!
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
*/

package main

import (
	"fmt"
)

func main() {

	// buffer 1, capaz de receber 1 elemento
	c := make(chan int, 1)

	// Função anônima auto-executavel
	func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
