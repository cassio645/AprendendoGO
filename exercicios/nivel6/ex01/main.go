/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/

package main

import "fmt"

func main() {
	total := soma(10, 5)
	fmt.Println(total)

	minhaIdade, voto := idadeVotar(1958)
	fmt.Printf("Você tem %d anos.\nSituação: %s\n", minhaIdade, voto)
}

// Função que retorna um int, soma dois valores e retorna o resultado
func soma(x int, y int) int {
	return x + y
}

// Função que retorna a idade e a situação se deve votar ou não
func idadeVotar(x int) (int, string) {

	var situacao string
	idade := 2023 - x

	if idade < 16 {
		situacao = "Não vota"
	} else if (idade > 15 && idade < 18) || idade >= 65 {
		situacao = "Voto opcional"
	} else {
		situacao = "Voto obrigatório"
	}
	return idade, situacao

}
