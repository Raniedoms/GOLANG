package main

import "fmt"

//você esta criando um tipo = será usuario
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
	telefone telefone
}

type endereco struct {
	rua string
	num uint8
}

type telefone struct {
	ddd    uint8
	number int
}

//coleção de campos

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Ranie"
	u.idade = 30
	fmt.Println(u)

	enderecoU2 := endereco{"Rua Itapera", 180}

	telefoneU2 := telefone{11, 996015525}
	u2 := usuario{"Tamires", 27, enderecoU2, telefoneU2}
	fmt.Println(u2)

	//Passando somente um valor.
	u3 := usuario{idade: 30}
	fmt.Println(u3)

}
