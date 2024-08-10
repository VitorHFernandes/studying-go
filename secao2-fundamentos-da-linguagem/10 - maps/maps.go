package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Vítor",
			"ultimo":   "Fernandes",
		},
		"curso": {
			"nome":         "Engenharia de Software",
			"universidade": "UniCesumar",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")

	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)

}
