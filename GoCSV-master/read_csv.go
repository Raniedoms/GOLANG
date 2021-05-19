package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// CSV : csv
type CSV struct { //estrutura que receberá os dados do CSV
	Request  string `json:"request"`
	Response string `json:"response"`
}

func checkErr(err error) { //checa erros
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func main() {
	//sem arquivo
	//var csvFile = strings.NewReader(`name;nickname;text
	//Valéria;Valchan;has been here!`)

	// com arquivo
	csvFile, err := os.Open("test.csv") //abre arquivo
	checkErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile)) //lê arquivo
	reader.Comma = ';'                                //define delimitador

	var person []CSV

	for {
		line, err := reader.Read() //para cada linha
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}
		person = append(person, CSV{ //adiciona uma pessoa
			Request:  line[0],
			Response: line[1],
		})
	}

	//personJSON, err := json.Marshal(person) //converte para JSON
	//checkErr(err)
	//fmt.Println(string(personJSON)) //exibe dados do csv
	//[{"name":"name","nickname":"nickname","text":"text"},{"name":"Valéria","nickname":"Valchan","text":"has been here!"}]
	fmt.Println("{{$root := initTag}}")
	fmt.Println("{{$addTag $root " + string(person[1].Request) + " " + person[1].Response + "}}") //exibe dados do csv
	// Valchan has been here!
}
