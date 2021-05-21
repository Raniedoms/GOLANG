package main

import "fmt"

//parametro por copia para essa funcao
func inverterSinal(numero int) int {
	return numero * -1
}

//referencia para essa funcao
//passando uma referencia para essa funcao, qlq alteracao feita vai impactar a variavel também fora da funcao
func inverterSinalcomPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)
	fmt.Println("----------------------------------------")
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalcomPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
