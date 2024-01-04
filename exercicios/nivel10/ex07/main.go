/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := make(chan int)

	// Primeiro for roda 10x e cria uma goroutine a cada iteração
	// cada goroutine envia 10 valores para o canal e espera 1segundo
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- rand.Intn(100) // rand.Int(100) valor inteiro aleatorio até 100
				time.Sleep(1 * time.Second) // espera 1s
			}
		}()

	}

	// Demonstra todos os valores recebidos no canal
	// Assim a cada 1s aparece mais 10 valores recebidos
	for valor := 0; valor < 100; valor++ {
		fmt.Println(valor, <-canal)
	}

}
