/*
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
*/

package main

import "fmt"

// Struct pai  Veiculo com portas e cor
type Veiculo struct {
	portas int
	cor    string
}

// Struct filho com traçãoNasQuatro sendo bool
type Caminhonete struct {
	Veiculo
	traçãoNasQuatro bool
}

// Struct filho com modeloLuxo sendo bool
type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func main() {

	// Criando o sedan luxo  (criado explicito)
	bmwM4 := Sedan{
		Veiculo:    Veiculo{portas: 4, cor: "Branco"},
		modeloLuxo: true,
	}

	// Criando a caminhonete com tração  (criado implicito)
	amarok := Caminhonete{Veiculo{4, "Preto"}, true,}

	// Demonstrando os valores criados
	fmt.Println("BMW M4\n\tPortas:", bmwM4.Veiculo.portas, "\n\tCor:", bmwM4.Veiculo.cor, "\n\tModelo Luxo:", bmwM4.modeloLuxo)
	fmt.Println("Amarok\n\tPortas:", amarok.portas, "\n\tCor:", amarok.cor, "\n\tTração nas 4 rodas:", amarok.traçãoNasQuatro)
}
