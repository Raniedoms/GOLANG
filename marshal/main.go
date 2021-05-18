package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type City struct {
	Name string `json:"city_name"`
	GDP  int    `json:"city_gpd"`
	//GDP        int    `json:"-"` DESSA FORMA IRÁ EXCLUIR O CAMPO
	Population int `json:"city_population"`
}

type User struct {
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	City      City       `json:"city"`
	CreatedAt customTime `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//omitempty - caso o campo seja vazio ele não enviará a tag
}

type customTime struct {
	time.Time
}

const layout = "2006-01-02"

func (c customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Format((layout)))), nil
}

func (c *customTime) UnmarshalJSON(v []byte) error {
	//"2001-01-01"
	var err error
	c.Time, err = time.Parse(layout, strings.ReplaceAll(string(v), "\"", ""))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	f, err := os.Open("out.Json")

	if err != nil {
		panic(err)
	}

	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	u := User{}
	if err := json.Unmarshal(jsonBytes, &u); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", u)

	////t := time.Now()
	//u := User{
	//	Name:      "Bob",
	//	Age:       20,
	//	City:      City{Name: "London", GDP: 500, Population: 800000},
	//	CreatedAt: customTime{time.Now()},
	//}
	//
	//out, err := json.Marshal(u)
	//if err != nil {
	//	log.Print(err)
	//}
	//
	//fmt.Println(string(out))

}
