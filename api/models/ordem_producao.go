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
	IDCliente	int64 `json:"IDCliente"`
	IDMaterial int64 `json:"IDMaterial"`
}

type OrdensDeProducao struct {
	Ordens map[int]OrdemProducao `json:"ordensDeProducao"`
}

func CreateTableOrdemProducao(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Clientes (
			IDCliente INTEGER PRIMARY KEY,
			Nome TEXT UNIQUE
		);
		CREATE TABLE IF NOT EXISTS Material (
			IDMaterial INTEGER PRIMARY KEY,
			CodigoMaterial TEXT UNIQUE,
			DescricaoMaterial TEXT UNIQUE
		);
		CREATE TABLE IF NOT EXISTS OrdemProducao (
			IDOrdemProducao INTEGER PRIMARY KEY,
			DataEntrega TEXT,
			CodigoOrdemProducao INTEGER UNIQUE,
			IDCliente INT,
			IDMaterial INT,
			Quantidade INTEGER,
			DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			IsAtivo BOOLEAN DEFAULT TRUE,
			FOREIGN KEY (IDCliente) REFERENCES Clientes(IDCliente),
			FOREIGN KEY (IDMaterial) REFERENCES Material(IDMaterial)
		);
	`)
	
	return err
}

func GetOrdemProducao(db *sql.DB) ([]OrdemProducao, error) {

	query := `SELECT IDOrdemProducao, 
		DataEntrega,
		CodigoOrdemProducao,
		IDCliente,
		IDMaterial,
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
			&o.IDCliente,
			&o.IDMaterial,
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

	insertClientQuery := `INSERT OR IGNORE INTO CLIENTES (Nome) VALUES (?)`

	res, err := db.Exec(insertClientQuery, o.Cliente)
	if err != nil {
		return err
	}

	IDCliente, err := res.LastInsertId()
	if err != nil {
		return err
	}

	insertMaterialQuery := `INSERT OR IGNORE INTO Material (CodigoMaterial, DescricaoMaterial) VALUES (?, ?)`

	res, err = db.Exec(insertMaterialQuery, o.CodigoMaterial, o.DescricaoMaterial)
	if err != nil {
		return err
	}

	var IDMaterial int64
	IDMaterial, err = res.LastInsertId()
	if err != nil {
		return err
	}

	insertQuery := `INSERT OR IGNORE INTO OrdemProducao (DataEntrega, CodigoOrdemProducao, IDCliente, IDMaterial, Quantidade) VALUES (?, ?, ?, ?, ?);`


	res, err = db.Exec(insertQuery, o.DataEntrega, o.CodigoOrdemProducao, IDCliente, IDMaterial, o.Quantidade)
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
