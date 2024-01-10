/*
	EX01
- Crie um package "cachorro".
    - Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
    - Documente seu código com comentários, e utilize a função Idade na sua função main.
- Rode seu programa para verificar se ele funciona.
- Rode um local server com godoc e leia sua documentação.
*/

/*
	EX02
- Coloque seu código no GitHub.
- Faça sua documentação aparecer em godoc.org.
*/

package main

import (
	"fmt"

	"github.com/cassio645/aprendendogo/exercicios/nivel12/ex01/cachorro"
)

func main() {
	anos := 10
	anosCaninos := cachorro.Idade(anos)

	fmt.Printf("Seu cachorro tem %v anos.\n", anosCaninos)
}
