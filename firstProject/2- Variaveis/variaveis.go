package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2" //declaração implicita (inferencia de tipo)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string  = "lalalala"
		variavel4 string  = "abcdef"
		variavel5 float64 = 1234
	)

	fmt.Println(variavel3, variavel4, variavel5)

	variavel6, variavel7 := "variavel 6", "variavel 7"
	fmt.Println(variavel6, variavel7)

	const constante1 string = "Constante 1" //não poderá alterar o valor
	fmt.Println(constante1)

	//invertendo valores de variaveis
	variavel6, variavel7 = variavel7, variavel6
	fmt.Println(variavel6, variavel7)

}
