package repositorio

import (
	"API/API/src/modelos"
	"database/sql"
)

//Usuarios representa um repositorio de usuários
type usuarios struct {
	//o usuarios vai ser minusculo - usar fiuncao para manipular essa funcao
	//struct que vai receber o que vier do banco, ele vai fazer a iteração
	db *sql.DB
}

//vai receber um banco que vai ser aberto pelo controller, e o controller que vai chamar esse repositorio de usuarios
//e essa funcao vai pegar esse banco e jogar dentro de um struct de um usuario - dentro desse struct vao ter os metodos de conexao com as tabelas do banco de dados.
// cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//criar - insere um usuario no banco de dados
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}