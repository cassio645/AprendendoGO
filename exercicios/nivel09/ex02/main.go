/*
	- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
package main

import "fmt"

// Struct do tipo pessoa com método falar
type Pessoa struct{
	nome string
	idade int
}
func (p *Pessoa) falar(){
	fmt.Println("Olá meu nome é", p.nome)
}

// Interface Humanos implementando o método falar
type Humanos interface{
	falar()
}
func dizerAlgumaCoisa(h Humanos){
	h.falar()
}

func main() {
	// Criando uma pessoa
	primeirapessoa := Pessoa{
		nome: "Cássio",
		idade: 26,
	}

	// Metodo dizerAlgumaCoisa recebendo o ponteiro de Pessoa
	dizerAlgumaCoisa(&primeirapessoa)

	// dizerAlgumaCoisa(primeirapessoa) - Pessoa does not implement Humanos (mthod falar has pointer receiver)
}
