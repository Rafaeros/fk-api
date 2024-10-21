package models

import (
	"fmt"
	"time"
	"database/sql"
)

type Material struct {
	IDMaterial int64 `id:"IDMaterial"`
	CodigoMaterial string `json:"codigoMaterial"`
	DescricaoMaterial string `json:"descricaoMaterial"`
	DataCriacao time.Time `json:"dataCriacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
	IsAtivo bool `json:"isAtivo"`
}

func CreateTableMaterial(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Material (
			IDMaterial INTEGER PRIMARY KEY,
			CodigoMaterial TEXT UNIQUE,
			DescricaoMaterial TEXT,
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

func CreateMaterial(db *sql.DB, CodigoMaterial string, DescricaoMaterial string) (int64, error) {
	var IDMaterial int64
	// Try to create Material
	insertQuery := `INSERT INTO Material (CodigoMaterial, DescricaoMaterial) VALUES (?, ?);`
	res, err := db.Exec(insertQuery, CodigoMaterial, DescricaoMaterial)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: Material.CodigoMaterial" {
			getQuery := `SELECT IDMaterial FROM Material WHERE CodigoMaterial = ?;`
			err := db.QueryRow(getQuery, CodigoMaterial).Scan(&IDMaterial)
			if err != nil {
			return 0, fmt.Errorf("erro ao obter ID do material existente: %v", err)
			}
		}
	} else {
		IDMaterial, err = res.LastInsertId()
		if err != nil {
		return 0, fmt.Errorf("erro ao obter ID do material inserido: %v", err)
		}
	}

	return	IDMaterial, nil
}