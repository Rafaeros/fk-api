package models

import (
	"fmt"
	"time"
	"database/sql"
)

type Cliente struct {
	IDCliente int64 `id:"IDCliente"`
	Nome string `json:"nome"`
	DataCriacao time.Time `json:"dataCriacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
	IsAtivo bool `json:"isAtivo"`
}

func CreateTableCliente(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Cliente (
			IDCliente INTEGER PRIMARY KEY,
			Nome TEXT UNIQUE,
			DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			IsAtivo BOOLEAN DEFAULT TRUE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func CreateCliente(db *sql.DB, NomeCliente string) (int64, error) {
	var IDCliente int64

	// Try to create Cliente
	insertQuery := `INSERT INTO CLIENTE (Nome) VALUES (?);`
	res, err := db.Exec(insertQuery, NomeCliente)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: CLIENTE.Nome" {
			getQuery := `SELECT IDCliente FROM CLIENTE WHERE Nome = ?;`
			err := db.QueryRow(getQuery, NomeCliente).Scan(&IDCliente)
			if err != nil {
			return 0, fmt.Errorf("erro ao obter ID do cliente existente: %v", err)
			}
		}
	} else {
		IDCliente, err = res.LastInsertId()
		if err != nil {
		return 0, fmt.Errorf("erro ao obter ID do cliente inserido: %v", err)
		}
	}

	return	IDCliente, nil
}