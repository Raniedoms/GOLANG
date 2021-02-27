package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	//map
	fmt.Println(usuario)
	fmt.Println("---------------------------------------")
	fmt.Println(usuario["nome"])

}
