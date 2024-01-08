/*
- Crie um struct "erroEspecial" que implemente a interface builtin.error. 
- Crie uma função que tenha um valor do tipo error como parâmetro. 
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/


package main

import "fmt"

// Struct implementando meu erro
type erroEspecial struct{
	meuErro string
}

// Função Error() do struct erroEspecial que cria uma mensagem de erro
func (e erroEspecial) Error() string{
	return fmt.Sprintln("Deu erro!", e.meuErro)
}

// Função que recebe um erro como parâmetro e printa
func deuErro(e error){
	fmt.Println("Recebendo um erro com a mensagem:\n", e)
}

func main() {
	arg := erroEspecial{"Alguma coisa errada aí."}
	deuErro(arg)

}
