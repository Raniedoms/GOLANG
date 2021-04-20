package main

import (
	"fmt"
	"time"
)

func main(){
	
	//i := 0
	//for i < 10 {
	//	fmt.Println("Incrementando i")
	//	//o time.Sleep da uma pausa de um segundo para voce verificar o que esta fazendo 
	//	time.Sleep(time.Second)
	//	i++
	//
	//	}

	//fmt.Println(i)

	
	//for j := 0; j < 10; j+= 2{
	//	fmt.Println("Incrementando j", j)
	//	time.Sleep(time.Second)
	//}

	//range vc iterar em um array, slice, string, etc
	nomes := []string {"Joao", "Davi", "Lucas"}
	//o primeiro e o indice que se refere a posicao
	//ai e o valor dele msm
	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	fmt.Println("---------------------")

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	fmt.Println("---------------------")

	for indice, letra := range "PALAVRA"{
		fmt.Println(string(letra))
		fmt.Println(indice)
	}

	fmt.Println("---------------------")

	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario{
		fmt.Println(chave,valor)
	}

	fmt.Println("---------------------")

	type usuarioStruct struct{
		nome string
		sobrenome string
	}


	//Loop infinito
	for {		
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second) 
	}
	
	//struct nao funciona range, somente em maps

	//usuario2 := usuarioStruct{"Ze", "Junior"}

	//for chave, valor := range usuario2{
	//	fmt.Println(chave,valor)
	//}
}