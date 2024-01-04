// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

package main

import "fmt"

func main() {
	// Criando canal que vai receber 100 valores
	canal100 := make(chan int)

	// goroutine chamando uma funcão que adiciona 100 valores
	go add100(canal100)

	// Função que retira e demonstra todos os valores de um canal
	retirar(canal100)

}

// Adicionando 100 valores no canal e fechando o canal
func add100(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

// Demonstrando todos valores de um canal
func retirar(c <-chan int) {
	for valor := range c {
		fmt.Println(valor)
	}

}
