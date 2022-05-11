package main

import (
	"context"
	"fmt"
	"os"

	uuid "github.com/google/uuid"

	pgx "github.com/jackc/pgx/v4"
)

func connectDb() (*pgx.Conn, error) {
	connstr := os.Getenv("DATATESTE_URL")
	if len(connstr) == 0 {
		err := fmt.Errorf("URL indisponivel ou inexistente")
		return nil, err
	}
	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		fmt.Errorf("sem acessso ao banco de dados [%v]: %v", connstr, err)

	}
	return conn, nil
}

var conn *pgx.Conn

func init() {
	var err error
	conn, err = connectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "impossivel estabelecer conexão %v\n", err)
		os.Exit(1)
	}
}

type contact struct {
	cod       uuid.UUID
	nome      string
	sobrenome string
	telefone  string
	cidade    string
}

func insertAcontact(c contact) error {
	var sql string = "INSERT INTO contato (cod, nome, sobrenome, telefone, cidade) VALUES($1, $2, $3, $4, $5)"

	_, err := conn.Exec(context.Background(), sql, c.cod, c.nome, c.sobrenome, c.telefone, c.cidade)
	if err != nil {
		err = fmt.Errorf("erro ao inserir os dados: %v", err)
		return err
	}
	return nil

}

func main() {
	defer conn.Close(context.Background())
	ct := contact{
		cod:       uuid.New(),
		nome:      "fulano",
		sobrenome: "beltrano",
		telefone:  "5525222-22",
		cidade:    "dubay",
	}
	err := insertAcontact(ct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "não foi inserir o contato: %v\n", err)
		os.Exit(2)
	}

}
