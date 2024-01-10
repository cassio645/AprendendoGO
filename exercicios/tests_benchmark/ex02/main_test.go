package main

import "testing"

type Testes struct {
	dados    []int
	resposta int
}

// TestSoma recebe um slice de Testes e testa a função soma comparando os resultados retornados com os resultados esperados.
func TestSoma(t *testing.T) {
	listaTestes := []Testes{
		{dados: []int{1, 2, 3}, resposta: 6},
		{[]int{10, 10, 10}, 30},
		{[]int{50, 10, -20}, 40},
	}

	for _, valor := range listaTestes {

		retornado := soma(valor.dados...)

		if retornado != valor.resposta {
			t.Error("Expecteddd:", valor.resposta, "Got:", retornado)
		}

	}

}