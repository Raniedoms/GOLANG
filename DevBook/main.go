package main

import (
	"API/API/src/config"
	"API/API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	//testando se o config esta funcionando
	fmt.Printf("Escutando na porta %d", config.Porta)
	//por padrao ele já via estar escutando no que esta configurado no .env
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

	//%d - para ficar no mesmo formato que tinha antes
}
