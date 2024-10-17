package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaeros/fk-api/api/connection"
	"github.com/rafaeros/fk-api/api/models"
)

func RoutersOrdemProducao(routers *mux.Router, base string) *mux.Router {
	routers.HandleFunc(base+"/ordem_producao/create", CreateTableOrdemProducaoHandler).Methods("GET")
	routers.HandleFunc(base+"/ordem_producao", GetOrdemProducaoHandler).Methods("GET")
	//routers.HandleFunc(base+"/ordem_producao/{id}", GetOrdemProducaoById).Methods("GET")
	routers.HandleFunc(base+"/ordem_producao", CreateOrdemProducaoHandler).Methods("POST")
	//routers.HandleFunc(base+"/ordem_producao/{id}", UpdateOrdemProducao).Methods("PUT")
	//routers.HandleFunc(base+"/ordem_producao/{id}", DeleteOrdemProducao).Methods("DELETE")
	return routers
}

func CreateTableOrdemProducaoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connection.OpenConnection()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return
	}
	defer db.Close()

	err = models.CreateTableOrdemProducao(db)
	if err != nil {
		fmt.Fprint(w, "Error creating table: ", err)
		return
	}

	fmt.Fprint(w, "Table OrdemProducao created successfully")
}

func GetOrdemProducaoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connection.OpenConnection()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		http.Error(w, "Error opening database connection", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	ordensDeProducao, err := models.GetOrdemProducao(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting OrdemProducao: %v", err), http.StatusInternalServerError)
		return
	}

	// convert to json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ordensDeProducao); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("OrdemProducao retrieved successfully")
}

func CreateOrdemProducaoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connection.OpenConnection()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		http.Error(w, "Error opening database connection", http.StatusInternalServerError)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error closing database connection: %v", err), http.StatusInternalServerError)
		}
	}()

	var o models.OrdemProducao
	err = json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading json: %v", err), http.StatusNotAcceptable)
		return
	}

	newOrdemProducao, err := models.CreateOrdemProducao(db, o)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating OrdemProducao: %v", err), http.StatusInternalServerError)
		return
	}

	// convert to json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOrdemProducao)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "OrdemProducao created successfully")
}
