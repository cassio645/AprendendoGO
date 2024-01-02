// - Utilize mutex para consertar a condição de corrida do exercício anterior.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador = 0

// criando um mutex
var mu sync.Mutex

const minhasGoroutines = 500

func main() {
	CriarGoroutines(minhasGoroutines)
	wg.Wait()

	// Usando mutex os valores são iguais
	fmt.Println("Contador:", contador, "\nGoroutines:", minhasGoroutines)
}

func CriarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {

		go func() {
			mu.Lock() // travando a função com mutex
			valor := contador
			runtime.Gosched()

			valor++
			contador = valor
			wg.Done()
			mu.Unlock() // destravando a função com mutex
		}()
	}
}
