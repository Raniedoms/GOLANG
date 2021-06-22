package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Todas as funcoes terao que ter esse formato - Rota representa todas as rotas da API

//Rotas representa todas as rotas da aplicação Web
type Rota struct {
	Uri    string
	Metodo string
	//como as funções que lidam com requisições http tem esse formato, todas as funções dentro da rota também terão que ter esse formato
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//retornar o router com todas as rotas configuradas
//ele vai receber um parametro de mux.Router e vai retornar um mux.Router
//ele recebe o router que nao vai ter nenhuma rota dentro e vai configurar todas essas rotas
//e devolver o router pronto
//Configurar colcoa todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	//dessa forma você esta iterando e vai dar um HandleFunc nela
	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
