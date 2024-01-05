package main

import "fmt"

func main() {
	contador := contar()
	outroContador := contar()

	fmt.Println(contador())
	fmt.Println(contador())
	fmt.Println(contador())
	fmt.Println(outroContador())
	fmt.Println(outroContador())
	fmt.Println(outroContador())
}

// Função que será recebida pela variável
func contar() func() int {
	x := 0
	// Função que irá incrementar em x toda vez que for chamada
	return func() int {
		x++
		return x
	}
}
