package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {

	schemaLoader := gojsonschema.NewReferenceLoader("file:/home/rdomsnunes/JsonSchema/validate/schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file:/home/rdomsnunes/JsonSchema/validate/document.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return
	}

	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("OK\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}
