FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

COPY .. .

RUN CGO_ENABLED=1 GOOS=linux go build -o /fk-api .

RUN go mod tidy
RUN go mod vendor

EXPOSE 8090

CMD [ "fk-api" ]