package main

import "fmt"

// como a funcao soma vai retornar um numero inteiro voce passa um int a mais

func soma (numeros ...int) int{
	//esse cara e um slice entao da para interar nele
	fmt.Println(numeros)

	total := 0

	for _, numero := range numeros{
		total += numero
	}
	return total
}


func escrever (texto string, numeros ...int){
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
}

func main (){
	totalDaSoma := soma (1,2,3,4,5,6,7,8,9,10,11)
	fmt.Println(totalDaSoma)


	escrever("Ola mundo", 10, 2, 3, 4, 5, 6 ,7)
}
