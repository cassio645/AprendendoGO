// - Crie um programa que demonstre o funcionamento da declaração if.

package main

import "fmt"

func main() {

	idade := 65

	// Menores de 16 anos não votam
	if idade < 16 {
		fmt.Println("Não pode votar!")

	} else if idade < 18 || idade > 64 { // 16, 17 OU pessoas mais de 64 anos votam por opcao
		fmt.Println("Votar é opcional!")

	} else {
		fmt.Println("Votar é obrigatório!") // caso contrário é obrigado a votar
	}
}
