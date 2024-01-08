/*
- Utilizando este código: https://play.golang.org/p/3W69TH4nON
- ...remova o underscore e verifique e lide com o erro de maneira apropriada.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// removendo _ e adicionando uma variável err para o possível erro
	bs, err := json.Marshal(p1.First)

	// Se o erro for diferente de nada, mostre o log do erro
	if err != nil{
		log.Println("Deu erro", err)
	}
	fmt.Println(string(bs))

}
