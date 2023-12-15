// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

func main() {
	
	for i := 10; i <= 100; i++ {

		// Calculando o resto da divisão por 4
		resto := i % 4
		
		fmt.Printf("%d %% 4 = %d\n", i, resto)
	}
}
