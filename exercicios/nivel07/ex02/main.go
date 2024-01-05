/*
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
*/

package main

import "fmt"

// struct Pessoa
type Pessoa struct {
	nome  string
	idade int
}

func main() {
	// Criando pessoa
	p1 := Pessoa{
		nome:  "Cássio",
		idade: 26,
	}

	fmt.Printf("Pessoa Criada\nNome: %s\nIdade: %d\n\n", p1.nome, p1.idade)

	mudeMe(&p1) // Chamando a função mudeMe, e passando o ponteiro da variável como argumento

	fmt.Println("Aniversário\nNova idade:", p1.idade) // Mostrando o valor após executar a função

}

// Função mudeMe que recebe como parâmetro um ponteiro para o tipo Pessoa
func mudeMe(pessoa *Pessoa) {

	// pessoa.idade++   // assim funciona também

	(*pessoa).idade++ // incrementando a idade da pessoa através do ponteiro
	
}
