package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaeros/fk-api/api/connection"
	"github.com/rafaeros/fk-api/api/models"
)

func RoutersCliente(routers *mux.Router, base string) *mux.Router {
	routers.HandleFunc(base+"/cliente/create", CreateClienteTableHandler).Methods("GET")
	//routers.HandleFunc(base+"/ordem_producao", GetOrdemProducaoHandler).Methods("GET")
	//routers.HandleFunc(base+"/cliente/{id}", GetOrdemProducaoById).Methods("GET")
	//routers.HandleFunc(base+"/cliente", CreateOrdemProducaoHandler).Methods("POST")
	//routers.HandleFunc(base+"/cliente/{id}", UpdateOrdemProducao).Methods("PUT")
	//routers.HandleFunc(base+"/cliente/{id}", DeleteOrdemProducao).Methods("DELETE")
	return routers
}

func CreateClienteTableHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connection.OpenConnection()
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return
	}
	defer db.Close()

	err = models.CreateTableCliente(db)
	if err != nil {
		fmt.Fprint(w, "Error creating table: ", err)
		return
	}

	fmt.Fprint(w, "Table Cliente created successfully")
}