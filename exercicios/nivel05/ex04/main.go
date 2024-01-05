/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import "fmt"

func main() {

	// Criando o struct anonimo e declarando os valores
	pessoa := struct {
		nome                  string
		linguagensProgramacao map[string]bool
		saboresSorvete        []string
	}{
		nome:                  "Cássio",
		linguagensProgramacao: map[string]bool{"Python": true, "Ruby": false, "TypeScript": false, "GO": true, "Java": true},
		saboresSorvete:        []string{"Cupuaçu", "Maracujá", "Napolitano"},
	}


	// Demonstrando os valores do struct anonimo
	fmt.Println(pessoa.nome, "\nSorvete:")
	for _, sabor := range pessoa.saboresSorvete {
		fmt.Println("\t-", sabor)
	}

	fmt.Println("Linguagens de Programação: ")
	for chave, valor := range pessoa.linguagensProgramacao {
		fmt.Println("\t", chave, "=", valor)
	}

}
