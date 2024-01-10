// Pacote calc realiza operações matemáticas de Soma e Multiplicação
package calc

// Soma recebe um número indeterminado de inteiros e retorna a soma deles.
func Soma(valores ...int) int {
	total := 0
	for _, valor := range valores {
		total += valor
	}
	return total
}

// Multiplica recebe um número indeterminado de inteiros e retorna a soma deles.
func Multiplica(valores ...int) int {
	total := 1
	for _, valor := range valores {
		total *= valor
	}
	return total
}
