# Etapa de build usando uma imagem oficial do Golang
FROM golang:1.23.1 AS builder

# Definir o diretório de trabalho no contêiner
WORKDIR /app

# Copiar o arquivo de dependências (go.mod e go.sum) para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências do projeto
RUN go mod download

# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar a aplicação
RUN go build -o fk-api .
# Etapa final - Imagem otimizada para produção (usando Alpine para reduzir o tamanho)
FROM alpine:latest

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário gerado na etapa de build
COPY --from=builder /app/fk-api .

# Expor a porta utilizada pela aplicação (se necessário)
EXPOSE 8080

# Definir o comando de inicialização padrão
CMD ["./fk-api"]