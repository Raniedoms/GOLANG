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
	log.Fatal(http.ListenAndServe(":8002", nil))

}

func indexPosHandlerRequest(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	r.ParseForm()

	request := r.PostForm.Get("request")
	transformation := r.PostForm.Get("transformation")

	newTemplate := template.New("todos")
	newTemplate = newTemplate.Funcs(sprig.FuncMap())

	t, err := newTemplate.Parse(transformation)
	if err != nil {
		panic(err)
	}

	payloadMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(request), &payloadMap)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var output bytes.Buffer
	err = t.Execute(&output, payloadMap)
	w.Write([]byte("request: " + request + "\n\t"))
	w.Write([]byte("transformation: " + transformation + "\n\t"))

	templates.ExecuteTemplate(w, "result.html", &output)
	//fmt.Print((output.Bytes()))

}

//
//data := make(map[string]interface{})
//err := json.Unmarshal([]byte(request), &data)
//jsonByte, err := json.Marshal(request)
//if err != nil {
//	panic(err)
//}
//fmt.Print(jsonByte)
//fmt.Print(err)
//fmt.Print(&data)
//
//t, err := template.New("todos").Funcs(sprig.FuncMap()).Parse(transformation)
//if err != nil {
//	panic(err)
//}
//
//var output bytes.Buffer
//err = t.Execute(&output, "todos")
//if err != nil {
//	panic(err)
//
//}
//
////fmt.Println(bytes.NewBuffer(request))
////fmt.Print(`"Request": " ` + request)
////fmt.Print(`"Transformation": " ` + transformation)
//fmt.Print(transformation)
//w.Write(output.Bytes())
//fmt.Print((output.Bytes()))

//var tmpl *template.Template
//	tmpl, *transform.GetError() = template.New(filepath.Base(session.Route.Source.Transform.Args)).Funcs(funcMap).ParseFiles(session.Route.Source.Transform.Args)
//
//tpl := template.Must(template.New(filepath.Base(request)).Funcs(sprig.FuncMap()).ParseFiles(transformation))

//fmt.Print(tpl)

//payloadMap := make(map[string]interface{})

//*transform.GetError() = tmpl.Execute(&output, payloadMap)
//*transform.GetError() = json.Unmarshal([]byte(*session.To.CommandResponse()), &payloadMap)

//func preTransfor() string {
//	var output bytes.Buffer
//	transform := indexPosHandlerRequest(request)
//	var tmpl *template.Template
//	tmpl, *transform.GetError() = template.New(filepath.Base()).Funcs(funcMap).ParseFiles(trans)
//
//}

//type Page struct {
//	Title string
//	Body  []byte
//}
//
//func (p *Page) save() error {
//	filename := p.Title + ".input"
//	return ioutil.content := "content"
//}
//
//func loadPage(title string) (*Page, error){
//	filename := title + ".input"
//	body, err := ioutil.ReadFile(filename)
//	if err != nil{
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}
//
//
//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, _ := loadPage(title)
//	renderTemplate(w, "view", p)
//}
//
//func editHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/edit/"):]
//	p, err := loadPage(title)
//	if err != nil {
//		p = &Page{Title: title}
//	}
//	renderTemplate(w, "edit", p)
//}
//
//func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//	t, _ := template.ParseFiles(tmpl + ".html")
//	t.Execute(w, p)
//}
//

//func main() {
//	http.HandleFunc("/view/", viewHandler)
//	http.HandleFunc("/edit/", editHandler)
//	http.HandleFunc("/save/", saveHandler)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	templates.ExecuteTemplate(w, "index.html", nil)
//r.ParseForm()
//request := r.PostForm.Get("input")
//transformation := r.PostForm.Get("transformation")

//}

//tpl := template.Must(template.New(filepath.Base()).Funcs(FuncMap()).ParseFiles(session.Route.Source.Transform.Args))
