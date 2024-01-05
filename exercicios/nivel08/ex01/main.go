/*
- Partindo do código abaixo, utilize marshal para transformar []user em JSON.
    - https://play.golang.org/p/U0jea43X55
- Atenção! Tem pegadinha aqui.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// Pegadinha: os atributos do struct 'user' estavam começando com letra minuscula.
	// Alterei para letras maiusculas para ser possivel passar para json
	usersJson := toJson(users...)
	fmt.Println(usersJson)

}

func toJson(users ...user) []string {
	// Transformando o slice de structs em slice de strings json
	var jsonslice []string

	for _, user := range users {
		myuser, err := json.Marshal(user)

		if err != nil {
			fmt.Println("error:", err)
		}

		jsonslice = append(jsonslice, string(myuser))
	}
	return jsonslice

}
