package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//ARRAY - lista de valores - array tem tamanho fixo
	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Posicao 1"
	fmt.Println(array2)

	array3 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array3)

	//ele nao deixa com o tamanho do array livre, mas sim com a quantidades que vc atribuiu a ele
	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	//Slice é baseado no array - flexibilidade de tamanhos
	//slice - fatia de um array
	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	//TypeOf - ela te devolve o tipo de uma variavel
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//no slice tem a função APPEND - adicionar um item no slice e retornar um slice novo com o item adicionado

	slice = append(slice, 18)
	fmt.Println(slice)

}
