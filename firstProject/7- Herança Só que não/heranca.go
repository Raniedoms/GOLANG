package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//vocẽ passa o nome da outra struct, mas não passa um tipo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Ranie", "Doms Nunes", 30, 169}
	fmt.Println(p1)

	e1 := estudante{p1, "TI", "Fiap"}
	fmt.Println(e1)

	//Pegando somente o nome
	fmt.Println(e1.nome)
}
