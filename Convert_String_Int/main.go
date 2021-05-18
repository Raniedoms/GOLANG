package main

import (
	"fmt"
	"strconv"
)

func main() {

	minhaIdadeString := "29"
	fmt.Printf("My Age: %s\n", minhaIdadeString)

	idadeInt, err := strconv.Atoi(minhaIdadeString)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Idade em int: %d\n", idadeInt)
	idadeInt += 1
	fmt.Printf("Nova idade: %d\n", idadeInt)

}
