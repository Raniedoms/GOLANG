package main

import (
	"errors"
	"fmt"
)

func main() {
	//diferença de bits
	//int8, int16, int32, int64

	var numero int16 = 100
	fmt.Println(numero)

	//ele já reclama direto, nem compila
	//var numero2 int16 = 1000000000
	//fmt.Println(numero2)

	//usa a arquitetura do computador como base
	var numero3 int = 10000
	fmt.Println(numero3)

	numero4 := 10000000000
	fmt.Println(numero4)

	//uint - unsygned int
	//não é suportado por conta do sinal
	//var numero2 uint32 = -100000
	//fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero5 rune = 123456
	fmt.Println(numero5)

	//BYTE = UINT8
	var numero6 byte = 123
	fmt.Println(numero6)

	//float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//STIRNGS
	//não tem char
	var str string = "TEXTOOOOOOOO"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	//Quando passar um char := 'B' e dar um print, ele vai pegar o número que esta na tabela Asp.

	char := 'B'
	fmt.Println(char)

	//para string será uma string vazia, para int ou float será valor 0
	var texto int
	fmt.Println(texto)

	//o valor zero do booleano é false
	var booleano1 bool = true
	fmt.Println(booleano1)

	//tipo de dado que serve como valor zero = resultado <nil>
	//pacote nativo do GO para erros
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
