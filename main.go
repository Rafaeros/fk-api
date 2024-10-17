package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"

	"github.com/rafaeros/fk-api/api/routers"
)

func main() {
	basePath := "/api/v1"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	log.Println("Get environment variables")

	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	primaryUrl := os.Getenv("TURSO_FK_URL")
	dbName := os.Getenv("TURSO_FK_DATABASE")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
        os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, libsql.WithAuthToken(authToken))
	if err != nil {
        fmt.Println("Error creating connector:", err)
        os.Exit(1)
    }
	defer connector.Close()

	db := sql.OpenDB(connector)
    defer db.Close()

	r := mux.NewRouter()
		
	log.Println("Setting up routes")
	r = routers.RoutersOrdemProducao(r, basePath)

	log.Println("Setting up handlers")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/'")
		w.Write([]byte("Hello World"))
	}).Methods(http.MethodGet)

	r.HandleFunc(basePath+"/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/ping'")
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	r.HandleFunc(basePath+"/version", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/version'")
		w.Write([]byte("1.0.0"))
	}).Methods(http.MethodGet)

	// Set up CORS
	log.Println("Setting up CORS")
	r.Use(mux.CORSMethodMiddleware(r))

	// Start the HTTP server on port 8090
	log.Println("Starting HTTP server on http://localhost:8090" + basePath)
	log.Fatal(http.ListenAndServe(":8090", r))
}
