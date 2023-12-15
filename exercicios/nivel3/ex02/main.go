// - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.

package main

import "fmt"

func main() {

	// For percorrendo o valor unicode das letras maiuculas(65-90)
	for i := 65; i <= 90; i++ {

		// Outro for para printar 3x cada
		for j := 0; j < 3; j++ {

			// Print do valor unicode e da letra maiuscula referente
			fmt.Printf("%U - %s\t", i, string(rune(i)))
		}

		// Vai pra linha de baixo assim que printa as 3x determinada letra.
		fmt.Println("")
	}
}
