package main

import "fmt"

// Função ExempleSoma faz o teste e aparece na documentação do código como Exemple
func ExampleSoma() {
	fmt.Println(Soma(10, 5, 5))
	fmt.Println(Soma(4, 4, 4))
	fmt.Println(Soma(10, -20, 4))
	// Output:
	// 20
	// 12
	// -6
}
