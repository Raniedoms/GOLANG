package main

import "fmt"

func main() {
	//dessa forma você fala para o GO que é uma função anonima
	//declara ela pra mim, mas assim que terminar de declarar já executa a função
	func(texto string) {
		fmt.Println(texto)
	}("Passando Parametro") //Passar o valor do parametro String que a função vai receber o valor

	fmt.Println("________________________________________________")

	//entao dessa forma ela vai retornar uma string
	func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto) //Sprintf - usado para concatenar informações
		//dessa forma não ira printar na tela pois esta dando um retorno, pois nao tem nada que esta armazenando esse retorno
	}("Passando Parametro") //Passar o valor do parametro String que a função vai receber o valor

	fmt.Println("________________________________________________")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10) //Sprintf - usado para concatenar informações
	}("Passando Parametro") //Passar o valor do parametro String que a função vai receber o valor
	fmt.Println(retorno)
}
