//- Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {
	hora := 0

	switch {
	case hora < 12 && hora >= 0:
		fmt.Println("Bom dia")
	case hora < 18 && hora >= 12:
		fmt.Println("Boa tarde")
	case hora >= 18 && hora <= 24:
		fmt.Println("Boa noite")
	default:
		fmt.Println("Hora inválida")
	}

}
