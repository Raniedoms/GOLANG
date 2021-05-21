package main

import "fmt"

// sao funcoes que referenciam variaveis que estao fora do seu corpo

//as funcoes podem retornar elas, passando como parametros etc, pois sao tipos normais
//ela tem uma lembranca de onde ela veio
func closure() func() {
	texto := "Dentro da função clousere"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
