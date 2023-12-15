/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu
*/

package main

import "fmt"

func main() {

	// Ano que nasci, valor inicial
	ano := 1997

	// Ano atual como valor final
	anoAtual := 2023

	// For com a condição de parada
	for ano <= anoAtual {

		fmt.Println(ano) // Printando o ano

		ano++ // incrementando a variável i
	}
}
