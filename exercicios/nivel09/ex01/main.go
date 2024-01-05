/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

// Declarando um waitgroup global para sincronizar as goroutines
var wg sync.WaitGroup

func main() {
	// Adicionando 2 ao waitgroup, pois teremos duas goroutines adicionais
	wg.Add(2)

	// Iniciando a primeira goroutine
	go PrimeiroPrint()

	// Iniciando a segunda goroutine
	go SegundoPrint()

	// Esperando até que todas as goroutines no waitgroup sejam concluídas
	wg.Wait()
}

func PrimeiroPrint() {
	fmt.Println("Função 1 Goroutines")

	// Sinalizando ao waitgroup que a goroutine foi concluída
	wg.Done()
}

func SegundoPrint() {
	fmt.Println("Função 2 Goroutines")

	// Sinalizando ao waitgroup que a goroutine foi concluída
	wg.Done()
}
