package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int64) (int64, int64) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	soma2 := somar(20, 40)
	fmt.Println(soma2)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(30, 20)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	//Quando você tem uma função com dois por exemplo e só quer pegar a soma por exemplo, na segunda vc coloca um _
	// A mesma coisa poderia ser feita para a subtração
	resultadosSoma1, _ := calculosMatematicos(20, 45)
	fmt.Println(resultadosSoma1)
}
