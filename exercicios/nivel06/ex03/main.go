// - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.


package main

import "fmt"

func main() {

	// Defer vai empilhar os comandos e executar por ultimo, desempilhando
	defer fmt.Println("Dezembro")
	defer fmt.Println("Novembro")

	fmt.Println("Janeiro")
	fmt.Println("Fevereiro")
	fmt.Println("Março")
	fmt.Println("Abril")

	defer fmt.Println("Outubro")

	fmt.Println("Maio")
	fmt.Println("Junho")
	fmt.Println("Julho")

	defer fmt.Println("Setembro")

	fmt.Println("Agosto")
	
/*
Janeiro
Fevereiro
Março
Abril
Maio
Junho
Julho
Agosto
Setembro
Outubro
Novembro
Dezembro
*/
	

}