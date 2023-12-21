/*
- Coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

package main

import "fmt"

// Struct Pessoa, com nome sobrenome e um slice para sabores de sorvete
type Pessoa struct {
	nome              string
	sobrenome         string
	sorvetesFavoritos []string
}

func main() {

	// Colocando os valores no map usando sobrenome como key
	pessoas := map[string]Pessoa{
		"Santos": {
			nome:              "Cássio",
			sobrenome:         "Santos",
			sorvetesFavoritos: []string{"Cupuaçu", "Napolitano", "Maracujá"},
		},

		"Rocha": {"Renato", "Rocha", []string{"Chocolate", "Morango", "Coco"}},
	}

	// Demonstrando os valores com range, e os sabores com outro range
	for chave, pessoa := range pessoas {
		fmt.Println("Chave:", chave, "Nome:", pessoa.nome)

		for _, sabor := range pessoa.sorvetesFavoritos {
			fmt.Println("\t-", sabor)
		}

		fmt.Println()
	}

}
