package main

import "fmt"

func recuperarExecucao() {
	//essa funcao recupera a funcao do nosso programa
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}
func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	//caso ele entre no panic, seu programa morre e nada do que seria executado após vai ser feito
	//para recuperar é o recover
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execucao!!!")
}
