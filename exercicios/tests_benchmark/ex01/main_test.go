package main

import "testing"

// TestSoma envia alguns números para a função soma e compara com o resultado esperado.
func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Expecteddd:", resultado, "Got:", teste)
	}
}
