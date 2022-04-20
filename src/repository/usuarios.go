package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) error {
	fmt.Println(usuario)
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values ($1, $2, $3, $4)",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if err != nil {
		return err
	}

	return nil
}
