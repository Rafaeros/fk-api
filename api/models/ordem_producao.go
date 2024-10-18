package models

import (
	"database/sql"
	"time"
)

type OrdemProducao struct {
	IDOrdemProducao     int64     `id:"IDOrdemProducao"`
	DataEntrega         string    `json:"dataEntrega"`
	CodigoOrdemProducao int32     `json:"codigoOrdemProducao"`
	Cliente             string    `json:"cliente"`
	CodigoMaterial      string    `json:"codigoMaterial"`
	DescricaoMaterial   string    `json:"descricaoMaterial"`
	Quantidade          int32     `json:"quantidade"`
	DataCriacao         time.Time `json:"dataCriacao"`
	DataAtualizacao     time.Time `json:"dataAtualizacao"`
	IsAtivo             bool      `json:"isAtivo"`
}

type OrdensDeProducao struct {
	Ordens map[int]OrdemProducao `json:"ordensDeProducao"`
}

func CreateTableOrdemProducao(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS OrdemProducao (
		IDOrdemProducao INTEGER PRIMARY KEY,
		DataEntrega TEXT,
		CodigoOrdemProducao INTEGER UNIQUE,
		Cliente TEXT,
		CodigoMaterial TEXT,
		DescricaoMaterial TEXT,
		Quantidade INTEGER,
		DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		IsAtivo BOOLEAN DEFAULT TRUE
		);
	`)

	return err
}

func GetOrdemProducao(db *sql.DB) ([]OrdemProducao, error) {

	query := `SELECT IDOrdemProducao, 
		DataEntrega,
		CodigoOrdemProducao,
		Cliente,
		CodigoMaterial,
	  	DescricaoMaterial,
	   	Quantidade,
	    DataCriacao,
		DataAtualizacao,
		IsAtivo FROM OrdemProducao`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordensDeProducao []OrdemProducao
	for rows.Next() {
		var o OrdemProducao
		err := rows.Scan(
			&o.IDOrdemProducao,
			&o.DataEntrega,
			&o.CodigoOrdemProducao,
			&o.Cliente,
			&o.CodigoMaterial,
			&o.DescricaoMaterial,
			&o.Quantidade,
			&o.DataCriacao,
			&o.DataAtualizacao,
			&o.IsAtivo,
		)
		if err != nil {
			return nil, err
		}

		ordensDeProducao = append(ordensDeProducao, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ordensDeProducao, nil
}

func (o *OrdemProducao) CreateOrdemProducao(db *sql.DB) error {

	insertQuery := `INSERT INTO OrdemProducao (DataEntrega, CodigoOrdemProducao, Cliente, CodigoMaterial, DescricaoMaterial, Quantidade) VALUES (?, ?, ?, ?, ?, ?);`


	res, err := db.Exec(insertQuery, o.DataEntrega, o.CodigoOrdemProducao, o.Cliente, o.CodigoMaterial, o.DescricaoMaterial, o.Quantidade)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	
	o.IDOrdemProducao = id
	return nil
}
