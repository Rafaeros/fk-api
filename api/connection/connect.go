package connection

import (
	"os"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func OpenConnection() (db *sql.DB, err error) {
	return sql.Open("sqlite3", "./fk.db")
}

func TursoConnect() (db *sql.DB, err error) {
	log.Println("Get environment variables")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	primaryUrl := os.Getenv("TURSO_FK_URL")
	dbName := os.Getenv("TURSO_FK_DATABASE")
	
	if authToken == "" || primaryUrl == "" || dbName == "" {
		log.Fatal("Missing environment variables")
		os.Exit(1)
	}

	url := fmt.Sprintf("libsql://%s-rafaeros.turso.io?auth_token=%s", dbName, authToken)

	log.Println("Setting up database connection")


	db, err = sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	fmt.Print("Created connection to turso")

	return db, nil
}