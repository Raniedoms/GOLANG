package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//criando como se fosse uma funcao, mas ela esta grudada a uma certa estrutura - usuario
//todos os usuarios tem o metodos salvar
// o u - variavel que pode usar para referenciar outros campos do mesmo usuario que chamou esse método
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)

}

//metodo de verificacao de maior de idade

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

//metodo fazer aniversario
//quando você tem um método que vai fazer alteração em algum campo do seu struct
//você passa a referencia do struct com o ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

	usuario2 := usuario{"Usuário 2", 30}
	fmt.Println(usuario2)
	usuario2.salvar()

	maiorIdade := usuario2.maiorIdade()
	fmt.Println(maiorIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
