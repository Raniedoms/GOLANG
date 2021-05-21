package main

import "fmt"

//pode ter uma funcao por pacote

var n int

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

//essa funcao entao executa primeiro /setar valor por exemplo ai utiliza a funcao init
func init() {
	fmt.Println("Executando a funcao init")
	n = 10
}
