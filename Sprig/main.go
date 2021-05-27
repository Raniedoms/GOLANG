package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	//r.HandleFunc("/transformation", indexHandler)
	r.HandleFunc("/transformation", indexPosHandlerRequest) //.Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func indexPosHandlerRequest(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	r.ParseForm()

	request := r.PostForm.Get("request")
	transformation := r.PostForm.Get("transformation")

	newTemplate := template.New("template")
	newTemplate = newTemplate.Funcs(sprig.FuncMap())

	t, err := newTemplate.Parse(transformation)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if err == nil {
		payloadMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(request), &payloadMap)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		var output bytes.Buffer
		err = t.Execute(&output, payloadMap)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte("request: " + request + "\n" + "transformation: " + transformation + "\n\t"))

		templates.ExecuteTemplate(w, "result.html", &output)
	}

	//fmt.Print((output.Bytes()))

}
