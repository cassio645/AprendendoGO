// - Crie uma variável de tipo string utilizando uma raw string literal.

package main

import "fmt"

func main() {
	rawString := `
	Minha string literal.
		Com várias linhas, e todo tipo de formatação sendo ignorada;
	\n \t %v %s %d "ping"
	`
	fmt.Println(rawString)
}
