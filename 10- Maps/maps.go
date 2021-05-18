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

	usuario2 := map[string]map[string]string{

		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "1",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"])
	fmt.Println("---------------------------------------")

	//quero apagar um map dentro desses
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	fmt.Println("---------------------------------------")

	//quero adicionar um map dentro desses
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)

}
