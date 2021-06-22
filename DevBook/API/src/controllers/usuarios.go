package controllers

import (
	"API/API/src/banco"
	"API/API/src/modelos"
	"API/API/src/repositorio"
	"API/API/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

//CriarUsuario - insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	//erro = errors.New("Deu Erro") - teste de erro
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//criando um usuario que esta no pacote de modelo, criar uma conexao com o banco de dados e passar essa conexao para um repositorio e através dos métodos insere as infos.
	var usuario modelos.Usuario
	//segundo parametro é o corpo de memória do usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return

	}
	defer db.Close()

	//passou o banco para dentro do repositório
	repositorio := repositorio.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

//BuscarUsuarios - busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)

	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)
}

//BuscarUsuario - busca um usuário no banco de dados

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	//os parametros vem no formato string
	//como id esta dentro deles e tem que ser compativel a struct uInt64
	//passa o usuarioId da rota na base 10 com 64 bits
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioDeUsuarios(db)

	usuario, erro := repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, usuario)
}

//AtualizarUsuario - atualiza um usuário no banco de dados

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado!"))
}

//DeletarUsuario - exclui um usuário no banco de dados

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário deletado!"))
}
