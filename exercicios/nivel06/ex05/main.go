/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

// Struct E método para o Quadrado
type Quadrado struct {
	lado float32
}

func (q Quadrado) area() {
	areaQuadrado := q.lado * q.lado
	fmt.Println("A área do quadrado é:", areaQuadrado)
}

// Struct E método para o Circulo
type Circulo struct {
	raio float64
}

func (c Circulo) area() {
	areaCirculo := math.Pi * (c.raio * c.raio)
	fmt.Println("A área do circulo é:", areaCirculo)
}

// Interface Figura com o método area()
type Figura interface {
	area()
}

// Método da interface recebe uma figura e chama o método area()
func info(f Figura) {
	f.area()
}

func main() {
	// Criando um quadrado
	quadrado := Quadrado{
		lado: 2.0,
	}

	// Criando um circulo
	circulo := Circulo{
		raio: 3.0,
	}

	// Demonstrando a área dos dois com o método da interface
	info(quadrado)
	info(circulo)

}
