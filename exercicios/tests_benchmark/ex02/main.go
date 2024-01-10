package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	fmt.Println(x)
}

// Função simples que soma varios inteiros e retorna o total
func soma(valores ...int) (total int) {
	total = 0
	for _, valor := range valores {
		total += valor
	}
	return
}
