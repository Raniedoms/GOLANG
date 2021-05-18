package controllers

import (
	"API/API/src/banco"
	"API/API/src/modelos"
	"API/API/src/repositorio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//CriarUsuario - insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal((erro))
	}

	//criando um usuario que esta no pacote de modelo, criar uma conexao com o banco de dados e passar essa conexao para um repositorio e através dos métodos insere as infos.
	var usuario modelos.Usuario
	//segundo parametro é o corpo de memória do usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)
}

//BuscarUsuarios - busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuário!"))
}

//BuscarUsuario - busca um usuário no banco de dados

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))

}

//AtualizarUsuario - atualiza um usuário no banco de dados

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado!"))
}

//DeletarUsuario - exclui um usuário no banco de dados

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário deletado!"))
}
