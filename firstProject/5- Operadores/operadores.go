package main

import "fmt"

func main() {
	//ARITMETICOS

	soma := 1 + 2

	subtracao := 1 - 2

	divisao := 10 / 4

	multiplicacao := 10 * 5

	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//se você for somar qualquer coisa que sejam de tipos diferentes não irá funcionar   int16 != int64
	//Por exemplo:
	//var numero1 int16 = 10
	//var numero2 int64 = 20
	//soma2 := numero1 + numero2
	//fmt.Println(soma2)

	//Dessa forma funciona e estou fazendo o soma2 com inferência
	var numero1 int16 = 10
	var numero2 int16 = 20
	soma2 := numero1 + numero2

	fmt.Println(soma2)

	//OPERADORES DE ATRIBUIÇÃO
	fmt.Println("-------OPERADORES DE ATRIBUIÇÃO-------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println("-------OPERADORES RELACIONAIS-------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	fmt.Println("-------OPERADORES LÓGICOS-------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS - incrementar o valor de uma variavel
	fmt.Println("-------OPERADORES UNÁRIOS-------")
	numero := 10
	numero++
	fmt.Println(numero)
	//somando + 10
	numero += 10
	fmt.Println(numero)

	numero3 := 15
	numero3 += 15 //numero = numero + 15
	fmt.Println(numero3)

	//DECREMENTAR

	numero3--
	fmt.Println(numero3)

	numero3 -= 20
	println(numero3)

	numero3 *= 3 //numero3 = numero3 * 3
	println(numero3)

	numero3 /= 2 //numero3 = numero3 / 2
	println(numero3)

}
