/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

package main

import "fmt"

// Struct Pessoa
type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// Método apresentar() para Pessoa
func (p Pessoa) apresentar() {
	fmt.Println("Olá meu nome é", p.nome, p.sobrenome, "e tenho", p.idade, "anos.")
}

func main() {
	// Criando uma pessoa
	primeiraPessoa := Pessoa{
		nome:      "Lucas",
		sobrenome: "da Silva",
		idade:     34,
	}

	// Utilizando o método apresentar()
	primeiraPessoa.apresentar()
}
