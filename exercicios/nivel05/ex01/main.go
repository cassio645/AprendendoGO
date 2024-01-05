/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
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

	// Criação do primeiro struct feito de forma explicita
	primeiraPessoa := Pessoa{
		nome:              "Cássio",
		sobrenome:         "Santos",
		sorvetesFavoritos: []string{"Cupuaçu", "Napolitano", "Maracujá"},
	}

	// Demonstrando os valores e um for para os sabores de sorvete
	fmt.Println("Nome:", primeiraPessoa.nome, primeiraPessoa.sobrenome)
	fmt.Println("Sabores de Sorvete Favorito:")
	for _, sabor := range primeiraPessoa.sorvetesFavoritos {
		fmt.Println("-", sabor)
	}

	fmt.Println() // pula linha

	// Criando o segundo struct de forma implicita
	segundaPessoa := Pessoa{"Renato", "Rocha", []string{"Chocolate", "Morango", "Coco"}}

	// Demonstrando os valores do segundo struct
	fmt.Println("Nome:", segundaPessoa.nome, segundaPessoa.sobrenome)
	fmt.Println("Sabores de Sorvete Favorito:")
	for _, sabor := range segundaPessoa.sorvetesFavoritos {
		fmt.Println("-", sabor)
	}

}
