package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJson)                  //! Retorna uma matriz de bytes
	fmt.Println(bytes.NewBuffer(cachorroJson)) //* Retorna o JSON esperado

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJson2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cachorroJson2))
}
