/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/

package main

import "fmt"

func main() {

	// Criando map com key e value como string
	dados := map[string][]string{
		"santos_cassio": {"programar", "assistir", "estudar"},
		"oliveira_joao": {"estudar", "jogar"},
		"korbes_ellen":  {"aula", "ensinar"},
	}

	// Demonstrando as keys e os values
	for chave, valor := range dados {
		fmt.Print(chave, "\n")
		for i, hobbie := range valor {
			fmt.Printf("\t %v - %v\n", i, hobbie)
		}
		fmt.Println()
	}
}
