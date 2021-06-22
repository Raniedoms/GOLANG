package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
)

var funcMap map[string]interface{} = template.FuncMap{
	"FormatNumber":  FormatNumber,
	"FloatToString": FloatToString,
	"MappingString": MappingString,
	"MappingFloat":  MappingFloat,
	"eqFloat":       eqFloat,
	"toString":      toString,
	"setTag":        setTag,
	"addTag":        addTag,
	"initTag":       initTag,
	"getPeriod":     getPeriod,
	"addf":          addf,
	"subf":          subf,
	"divf":          divf,
	"mulf":          mulf,
	"parseFloat":    parseFloat,
	"makeList":      makeList,
}

func init() {
	for key, value := range sprig.FuncMap() {
		funcMap[key] = value
	}
}

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/transformation", indexHandler)
	r.HandleFunc("/transformation", indexPosHandlerRequest) //.Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func indexPosHandlerRequest(w http.ResponseWriter, r *http.Request) {
	//estou transformando uma interface de map na variavel payloadMap
	payloadMap := make(map[string]interface{})
	//templates.ExecuteTemplate(w, "index.html", payloadMap)

	//Para todas as solicitações, o ParseForm analisa a consulta bruta do URL e atualiza o r.Form.
	r.ParseForm()

	//Pego as infos pelos Id's no form no index.html
	request := r.PostForm.Get("request")
	transformation := r.PostForm.Get("transformation")

	newTemplate := template.New("template")

	//executo um map de funcMap onde contem todas as funções que posso introduzir como também as funções do Mastermind/sprig
	newTemplate = newTemplate.Funcs(funcMap)

	payloadMap["transformation"] = transformation
	payloadMap["input"] = request

	t, err := newTemplate.Parse(transformation)
	if err != nil {
		payloadMap["result"] = err.Error()
	}

	if err == nil {
		//err = json.Unmarshal([]byte(request), &payloadMap)

		if err != nil {
			payloadMap["result"] = err.Error()
		}

		payloadTeste := make(map[string]interface{})
		err = json.Unmarshal([]byte(request), &payloadTeste)
		if err != nil {
			payloadMap["result"] = err.Error()
		}

		var output bytes.Buffer
		err = t.Execute(&output, payloadTeste)
		if err != nil {
			payloadMap["result"] = err.Error()
		}

		out := output.Bytes()

		var outputFormatado bytes.Buffer
		json.Indent(&outputFormatado, out, "", "	")

		payloadMap["result"] = &outputFormatado
		//w.Write([]byte("request: " + request + "\n" + "transformation: " + transformation + "\n\t"))
		if err != nil {
			payloadMap["result"] = err.Error()
		}
		fmt.Println(payloadMap)

	}

	t2, err2 := template.ParseFiles("templates/index.html")
	if err2 != nil {
		w.Write([]byte(err.Error()))
		payloadMap["result"] = err.Error()
	}

	t2.Execute(w, payloadMap)

	//fmt.Print((output.Bytes()))

}

func makeList(el interface{}) []interface{} {
	if el == nil {
		return make([]interface{}, 0)
	}
	switch el.(type) {
	case []interface{}:
		return el.([]interface{})
	default:
		l := make([]interface{}, 1, 1)
		l[0] = el
		return l
	}
}

func getItemFromListByIndex(list interface{}, index int) interface{} {
	switch list.(type) {
	case []interface{}:
		for i, j := range list.([]interface{}) {
			item := j.(map[string]interface{})
			if i == index {
				return item
			}
		}
	}
	return nil
}

func parseFloat(xs ...interface{}) (fs []float64, err error) {
	fs = make([]float64, len(xs))
	for i, x := range xs {
		switch v := x.(type) {
		case int:
			fs[i] = float64(v)
		case string:
			xf, err := strconv.ParseFloat(v, 64)
			if err != nil {
				errS := fmt.Sprintf("An error occurred while converting b: %s to float: %v", v, err)
				fmt.Println(errS)
				return fs, errors.New(errS)
			}
			fs[i] = xf
		case float64:
			fs[i] = v
		case float32:
			fs[i] = float64(v)
		default:
			errS := fmt.Sprintf("I don't know about type %T!\n", v)
			fmt.Println(errS)
			return fs, errors.New(errS)
		}
	}
	return fs, nil
}

func addf(a, b interface{}) float64 {
	fs, err := parseFloat(a, b)
	if err != nil {
		return 0
	}
	return fs[0] + fs[1]
}

func subf(a, b interface{}) float64 {
	fs, err := parseFloat(a, b)
	if err != nil {
		return 0
	}
	return fs[0] - fs[1]
}

func divf(a, b interface{}) float64 {
	fs, err := parseFloat(a, b)
	if err != nil {
		return 0
	}
	return fs[0] / fs[1]
}

func mulf(a, b interface{}) float64 {
	fs, err := parseFloat(a, b)
	if err != nil {
		return 0
	}
	return fs[0] * fs[1]
}

func getPeriod(layout, dateInitial, dateFinal string) int64 {
	initialDate, _ := time.Parse(layout, dateInitial)
	finalDate, _ := time.Parse(layout, dateFinal)
	duration := finalDate.Sub(initialDate)
	days := duration.Hours() / 24
	return int64(days)
}

func getFuncs() string {
	var funcs bytes.Buffer
	for k, _ := range funcMap {
		funcs.WriteString(k)
		funcs.WriteString(";")
	}
	return funcs.String()
}

func FormatNumber(value float64) string {
	return fmt.Sprintf("%.4f", value)
}

func FloatToString(value float64) string {
	return fmt.Sprintf("%f", value)
}

func MappingString(value interface{}) interface{} {
	if value == nil {
		switch value.(type) {
		case uint64:
			return 0
		case bool:
			return false
		default:
			return ""
		}
	} else {
		return value
	}
}

func MappingFloat(value interface{}) float64 {
	if value == nil {
		return 0
	} else {
		return value.(float64)
	}
}

func eqFloat(value1, value2 float64) bool {
	return value1 == value2
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch s := v.(type) {
	case string:
		return s
	default:
		f := fmt.Sprintf("%v", s)
		return f
	}
}

func addTag(m map[string]interface{}, k string, v interface{}) string {
	if v == nil {
		return ""
	}
	s := toString(v)
	if len(s) == 0 {
		return ""
	}

	setTag(m, k, v)
	return ""
}

func setTag(m map[string]interface{}, k string, v interface{}) string {
	if !strings.Contains(k, ".") {
		m[k] = v
		return ""
	}

	parents := strings.Split(k, ".")
	children := strings.Join(parents[1:], ".")
	if m[parents[0]] == nil {
		m[parents[0]] = make(map[string]interface{}, 1)
	}

	switch child := m[parents[0]].(type) {
	case map[string]interface{}:
		setTag(child, children, v)
		return ""
	default:
		m[parents[0]] = make(map[string]interface{}, 1)
		setTag(m[parents[0]].(map[string]interface{}), children, v)
		return ""
	}
}

func initTag() map[string]interface{} {
	return make(map[string]interface{})
}
