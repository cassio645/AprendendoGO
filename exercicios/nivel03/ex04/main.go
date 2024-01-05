/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import "fmt"

func main() {

	// Ano que nasci, valor inicial
	ano := 1997

	// Ano atual como valor final
	anoAtual := 2023

	// for..ever
	for {

		fmt.Println(ano) // Printando o ano

		ano++ // incrementando a variável i

		// Condição de parada com break
		if ano > anoAtual {
			break
		}
	}
}
