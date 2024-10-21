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
		FOREIGN KEY (IDCliente) REFERENCES Cliente(IDCliente)
);