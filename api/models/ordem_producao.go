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
	IDCliente           int64     `id:"IDCliente"`
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
		CodigoMaterial TEXT,
		DescricaoMaterial TEXT,
		Quantidade INTEGER,
		DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		IsAtivo BOOLEAN DEFAULT TRUE,
		IDCliente INTEGER,
		FOREIGN KEY (IDCliente) REFERENCES Cliente(IDCliente)
		);

		CREATE TABLE IF NOT EXISTS CLIENTE (
		IDCliente INTEGER PRIMARY KEY,
		Nome TEXT UNIQUE
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func GetOrdemProducao(db *sql.DB) ([]OrdemProducao, error) {

	query := `SELECT 
		IDOrdemProducao, 
		DataEntrega,
		CodigoOrdemProducao,
		CodigoMaterial,
	  	DescricaoMaterial,
	   	Quantidade,
		DataCriacao,
		DataAtualizacao,
		IsAtivo,
		IDCliente FROM OrdemProducao`

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
			&o.CodigoMaterial,
			&o.DescricaoMaterial,
			&o.Quantidade,
			&o.DataCriacao,
			&o.DataAtualizacao,
			&o.IsAtivo,
			&o.IDCliente,
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

	// Get or Create IDCLIENTE
	IDCliente, err := CreateCliente(db, o.Cliente)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO OrdemProducao (DataEntrega, CodigoOrdemProducao, CodigoMaterial, DescricaoMaterial, Quantidade, IDCliente) VALUES (?, ?, ?, ?, ?, ?);`

	res, err := db.Exec(insertQuery, o.DataEntrega, o.CodigoOrdemProducao, o.CodigoMaterial, o.DescricaoMaterial, o.Quantidade, IDCliente)
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
