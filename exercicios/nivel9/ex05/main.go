// - Utilize atomic para consertar a condição de corrida do exercício #3.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var contador int32

const minhasGoroutines = 500

func main() {
	CriarGoroutines(minhasGoroutines)

	wg.Wait()

	fmt.Println("Contador:", contador, "\nGoroutines:", minhasGoroutines)
}

func CriarGoroutines(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {

		go func() {
			// Usando atomic.AddInt32 para adicionar 1 ao contador através do endereço de memória
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()

			// Lendo atomicamente o valor da variável contador através do endereço de memoria
			atomic.LoadInt32(&contador)
			fmt.Println(contador)
			wg.Done()
		}()
	}
}
