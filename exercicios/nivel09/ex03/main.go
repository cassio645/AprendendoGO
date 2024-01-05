/*
- Utilizando goroutines, crie um programa incrementador:
  - Tenha uma variável com o valor da contagem
  - Crie um monte de goroutines, onde cada uma deve:
  - Ler o valor do contador
  - Salvar este valor
  - Fazer yield da thread com runtime.Gosched()
  - Incrementar o valor salvo
  - Copiar o novo valor para a variável inicial
  - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
  - Demonstre que há uma condição de corrida utilizando a flag -race
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// criando um waitgroup
var wg sync.WaitGroup

// criando uma variavel contador começando em zero e uma constante com o numero de goroutines
var contador = 0

const minhasGoroutines = 500

func main() {
	// Chamando a funcao que cria goroutines passando a constante como argumento
	CriarGoroutines(minhasGoroutines)

	// wait para o programa esperar as goroutines acabarem para finalizar
	wg.Wait()

	// Mostrando o valor do contador e a quantidade de goroutines criada (deveria ser o mesmo valor)
	fmt.Println("Contador:", contador, "\nGoroutines:", minhasGoroutines)
}

func CriarGoroutines(i int) {
	// adicionando o numero de goroutines
	wg.Add(i)

	for j := 0; j < i; j++ {

		go func() {
			// implementando +1 no contador a cada goroutine
			valor := contador
			runtime.Gosched()

			valor++
			contador = valor
			wg.Done()
		}()
	}
}
