package calc

import "testing"

type Testes struct {
	dados    []int
	resposta int
}

// Benchmark simples para medir o desempenho da função Soma calculando 2+2
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(2, 2)
	}
}

// Benchmark para medir o desempenho da função Soma usando um slice do struct Testes
func BenchmarkSomaStruct(b *testing.B) {

	// Criando um slice de testes
	listaTestes := []Testes{
		{dados: []int{1, 2, 3}, resposta: 6},
		{[]int{10, 10, 10}, 30},
		{[]int{50, 10, -20}, 40},
	}

	for i := 0; i < b.N; i++ {
		for _, valor := range listaTestes {
			Soma(valor.dados...)
		}
	}
}

// Benchmark calculando o desempenho da função Multiplica calculando 5*5
func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(5, 5)
	}
}


// 		EX06 - COVER: go test -cover    retorna a % de cobertura dos testes do modulo
// Criar testes aqui

// Testes da função Soma
func TestSoma(t *testing.T) {
	listaTestes := []Testes{
		{dados: []int{1, 2, 3}, resposta: 6},
		{[]int{10, 10, 10}, 30},
		{[]int{50, 10, -20}, 40},
	}

	for _, valor := range listaTestes {

		retornado := Soma(valor.dados...)

		if retornado != valor.resposta {
			t.Error("Expected:", valor.resposta, "Got:", retornado)
		}
	}
}

// Testes da função Multiplica
func TestMultiplica(t *testing.T) {

	listaTestes := []Testes{
		{dados: []int{3, 3}, resposta: 9},
		{[]int{10, 10, 10}, 1000},
		{[]int{5, 5, 10}, 250},
	}

	for _, valor := range listaTestes {

		retornado := Multiplica(valor.dados...)

		if retornado != valor.resposta {
			t.Error("Expected:", valor.resposta, "Got:", retornado)
		}
	}
}

// Gerar um arquivo cover:  go test -coverprofile [nome].out
// Abre uma página html do resultado cover: go tool cover -html=[nome].out