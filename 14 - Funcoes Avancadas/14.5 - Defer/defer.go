package main

import "fmt"

func funcao1() {
	fmt.Println("executando a funcao 1")
}

func funcao2() {
	fmt.Println("executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	//mesmo ele estando na primeira linha da função, ele só vai ser executado imediatamente depois de qualquer return
	defer fmt.Println("Média calculada. Resultado será retornado")

	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

//o defer quer dizer adiado - você fala para adiar a execução da função até o ultimo momento possível
func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(4, 8))
}
