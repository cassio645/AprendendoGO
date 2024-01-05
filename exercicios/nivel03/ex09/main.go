// - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Seu esporte favorito é Futebol")
	case "voleibol":
		fmt.Println("Seu esporte favorito é Voleibol")
	case "basquetebol":
		fmt.Println("Seu esporte favorito é Basquetebol")
	case "natação":
		fmt.Println("Seu esporte favorito é Natação")
	case "corrida":
		fmt.Println("Seu esporte favorito é Corrrida")

	}
}
