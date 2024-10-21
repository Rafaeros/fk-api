package models

import (
	"time"
	"database/sql"
)

type OrdemProducao struct {
	IDOrdemProducao     int64     `id:"IDOrdemProducao"`
	DataEntrega         string    `json:"dataEntrega"`
	CodigoOrdemProducao int32     `json:"codigoOrdemProducao"`
	Cliente             string    `json:"cliente"`
	CodigoMaterial      string    `json:"codigoMaterial"`
	DescricaoMaterial   string    `json:"descricaoMaterial"`
	Quantidade          int32     `json:"quantidade"`
	IDCliente           int64     `json:"IDCliente"`
	IDMaterial			int64	  `json:"IDMaterial"`
	DataCriacao         time.Time `json:"dataCriacao"`
	DataAtualizacao     time.Time `json:"dataAtualizacao"`
	IsAtivo             bool      `json:"isAtivo"`
}

type OrdensDeProducao struct {
	Ordens map[int]OrdemProducao `json:"ordensDeProducao"`
}

type OrdemProducaoResponse struct {
	IDOrdemProducao int64 `json:"IDOrdemProducao"`
	DataEntrega         string    `json:"dataEntrega"`
	CodigoOrdemProducao int32     `json:"codigoOrdemProducao"`
	Quantidade          int32     `json:"quantidade"`
	IDCliente           int64     `json:"IDCliente"`
	IDMaterial			int64	  `json:"IDMaterial"`
	IsAtivo             bool      `json:"isAtivo"`
}

func CreateTableOrdemProducao(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS OrdemProducao (
			IDOrdemProducao INTEGER PRIMARY KEY,
			DataEntrega TEXT,
			CodigoOrdemProducao INTEGER UNIQUE,
			Quantidade INTEGER,
			IDCliente INTEGER,
			IDMaterial INTEGER,
			DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			IsAtivo BOOLEAN DEFAULT TRUE,
			FOREIGN KEY (IDCliente) REFERENCES Cliente(IDCliente),
			FOREIGN KEY (IDMaterial) REFERENCES Cliente(IDMaterial)
		);

		CREATE TABLE IF NOT EXISTS Cliente (
			IDCliente INTEGER PRIMARY KEY,
			Nome TEXT UNIQUE
		);

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

func GetOrdemProducao(db *sql.DB) ([]OrdemProducaoResponse, error) {

	query := `SELECT 
		IDOrdemProducao, 
		DataEntrega,
		CodigoOrdemProducao,
	   	Quantidade,
		IDCliente,
		IDMaterial,
		IsAtivo FROM OrdemProducao`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordensDeProducao []OrdemProducaoResponse
	for rows.Next() {
		var o OrdemProducaoResponse
		err := rows.Scan(
			&o.IDOrdemProducao,
			&o.DataEntrega,
			&o.CodigoOrdemProducao,
			&o.Quantidade,
			&o.IDCliente,
			&o.IDMaterial,
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

	// Get or Create IDCLIENTE
	IDCliente, err := CreateCliente(db, o.Cliente)
	if err != nil {
		return err
	}

	IDMaterial, err := CreateMaterial(db, o.CodigoMaterial, o.DescricaoMaterial)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO OrdemProducao (DataEntrega, CodigoOrdemProducao, Quantidade, IDCliente, IDMaterial) VALUES (?, ?, ?, ?, ?);`

	res, err := db.Exec(insertQuery, o.DataEntrega, o.CodigoOrdemProducao, o.Quantidade, IDCliente, IDMaterial)
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
