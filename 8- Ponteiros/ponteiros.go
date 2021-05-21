package main

import "fmt"

//ponteiro - salva o endereço de memória de alguma coisa

func main() {
	fmt.Println("PONTEIROS")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	//quando voce atribui um valor a uma variavel, esse valor seria uma cópia.
	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int

	//O ponteiro guarda o endereço de memória de um valor inteiro por exemplo
	variavel3 = 100
	ponteiro = &variavel3 //dentro do endereço de memória esta o valor 10 - 0xc0000140e0

	fmt.Println(variavel3, ponteiro)

	//para ver o valor tem que colocar o * na frente dele
	//desreferenciação
	fmt.Println(variavel3, *ponteiro)

	//como o valor que esta armazenado no endereço de memória é diferente, o ponteiro muda seu valor dentro da variável.
	variavel3 = 150
	fmt.Println(ponteiro)
	fmt.Println(variavel3, *ponteiro)

}
