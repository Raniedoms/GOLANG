package main


import (
	"log"
	"os"
	"linha-de-comando/app"
	"fmt"
)



func main (){

	fmt.Println("Ponto de Partida")
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
}