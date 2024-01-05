/*
- Crie uma slice usando make que possa conter todos os estados do Brasil.
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice sem utilizar range.
*/

package main

import "fmt"

func main() {
	// Usando make para criar um slice que possa conter todos os estaods do Brasil
	estados := make([]string, 26, 26)

	// Adicionando os estados
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	// Demonstrando o len e o cap do slice
	fmt.Println("len:", len(estados), "cap:", cap(estados))

	// Demonstrando todos os estados
	for indice := 0; indice < len(estados); indice++ {
		fmt.Println(estados[indice])
	}
}
