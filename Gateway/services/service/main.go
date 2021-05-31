package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("CONTENT"))
}

func main() {
	//Quando acessar a URL, e baseado em uma porta na variavel de ambiente ele vai mostrar o resultado
	//que esta na função handler
	http.HandleFunc("/", handler)
	log.Fatal((http.ListenAndServe(os.Getenv("PORT"), nil)))
}
