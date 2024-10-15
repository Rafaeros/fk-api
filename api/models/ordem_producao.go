package models

import (
	"fmt"
	"time"
	"database/sql"
	"github.com/Rafaeros/fk-api/api/connection"
)

type OrdemProducao struct {
	IDOrdemProducao     int64
	DataEntrega         string `json:"dataEntrega"`
	CodigoOrdemProducao int32  `json:"codigoOrdemProducao"`
	Cliente             string `json:"cliente"`
	CodigoMaterial      string `json:"codigoMaterial"`
	DescricaoMaterial   string `json:"descricaoMaterial"`
	Quantidade          int32  `json:"quantidade"`
	DataCriacao         time.Time
	DataAtualizacao     time.Time
	IsAtivo             bool
}

func CreateTableOrdemProducao(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS OrdemProducao (
		IDOrdemProducao INTEGER PRIMARY KEY,
		DataEntrega TEXT,
		CodigoOrdemProducao INTEGER,
		Cliente TEXT,
		CodigoMaterial TEXT,
		DescricaoMaterial TEXT,
		Quantidade INTEGER,
		DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		IsAtivo BOOLEAN DEFAULT TRUE
	`)
	return err
}

func CreateOrdemProducao() {
	db, err := connection.OpenConnection()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()

}
