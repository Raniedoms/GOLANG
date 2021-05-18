package modelos

import "time"

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
