package modelos

import (
	"errors"
	"strings"
	"time"
)

//usuario representa um usuario utilizando a rede social
type Usuario struct {
	//quando for passar esse usuario para um json, e estiver em branco ele tira o campo do JSON
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

//Preparar vai chamar os métodos validar e formatar o usuário recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome nao pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nick nao pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email nao pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("A Senha nao pode estar em branco")
	}
	return nil
}

//retirar os espaços a mais com o TrimSpace
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
