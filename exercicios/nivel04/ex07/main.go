/*
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main

import "fmt"

func main() {

	// Criando slice de slice
	dados := [][]string{
		{"Cassio", "Santos", "Programar"},
		{"João", "Oliveira", "Estudar"},
		{"Ellen", "Korbes", "Aula"},
	}

	// Demonstrando os dados
	for pessoa := range dados {
		fmt.Printf("%v %v\t %v\n", dados[pessoa][0], dados[pessoa][1], dados[pessoa][2])
	}
}
