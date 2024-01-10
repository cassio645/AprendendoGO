package main

import "fmt"

func main() {
	x := Soma(1, 2, 3)
	fmt.Println(x)
}

// Função simples que soma varios inteiros e retorna o total
func Soma(valores ...int) (total int) {
	total = 0
	for _, valor := range valores {
		total += valor
	}
	return
}
