// Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.

package main

import "fmt"

func main() {

	for i := 33; i < 123; i++ {

		// Demonstrando decimal e a string referente àquele número.
		fmt.Printf("%d - %s\n", i, string(rune(i)))

	}
}
