//- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {

	// Criando map com key e value como string
	dados := map[string][]string{
		"santos_cassio": {"programar", "assistir", "estudar"},
		"oliveira_joao": {"estudar", "jogar"},
		"korbes_ellen":  {"aula", "ensinar"},
	}

	// Adicionando uma nova entrada no map
	dados["pike_rob"] = []string{"criar linguagem de programação", "usar termos loucos"}

	// Demonstrando as keys e os values
	for chave, valor := range dados {
		fmt.Print(chave, "\n")
		for i, hobbie := range valor {
			fmt.Printf("\t %v - %v\n", i, hobbie)
		}
		fmt.Println()
	}
}
