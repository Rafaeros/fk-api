CREATE TABLE IF NOT EXISTS Cliente (
    IDCliente INTEGER PRIMARY KEY,
    Nome TEXT UNIQUE,
    DataCriacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    DataAtualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    IsAtivo BOOLEAN DEFAULT TRUE
);