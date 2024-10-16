package routers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rafaeros/fk-api/api/models"
	"github.com/rafaeros/fk-api/api/connection"
)

func RoutersOrdemProducao(routers *mux.Router, base string) *mux.Router {
	routers.HandleFunc(base+"/ordem_producao/create", CreateTableOrdemProducaoHandler).Methods("GET")
	//routers.HandleFunc(base+"/ordem-producao", GetOrdemProducao).Methods("GET")
	//routers.HandleFunc(base+"/ordem-producao/{id}", GetOrdemProducaoById).Methods("GET")
	//routers.HandleFunc(base+"/ordem-producao", CreateOrdemProducao).Methods("POST")
	//routers.HandleFunc(base+"/ordem-producao/{id}", UpdateOrdemProducao).Methods("PUT")
	//routers.HandleFunc(base+"/ordem-producao/{id}", DeleteOrdemProducao).Methods("DELETE")
	return routers
}

func CreateTableOrdemProducaoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connection.OpenConnection()
	if err != nil {
		fmt.Fprint(w, "Error connecting to database:", err)
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

func CreateOrdemProducaoHandler(w http.ResponseWriter, r *http.Request) {
	
}