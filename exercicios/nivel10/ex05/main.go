/*
- Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
- ...demonstre o comma ok idiom.
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Go func executando de maneira concorrente
	go func() {
		c <- 42
		close(c)
	}()

	// Verificando o valor que esta no canal C e se esta ok
	v, ok := <-c
	fmt.Println(v, ok)

	// Verificando novamente, agora sem nenhum valor e ok dando false
	v, ok = <-c
	fmt.Println(v, ok)
}
