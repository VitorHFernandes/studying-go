package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"-"` //TODO => ignora a chave
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroJson), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	cachorroJson2 := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(cachorroJson2), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

}
