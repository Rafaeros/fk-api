package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaeros/fk-api/api/routers"
)

func main() {
	basePath := "/api/v1"
	r := mux.NewRouter()
	log.Println("Setting up routes")
	r = routers.RoutersOrdemProducao(r, basePath)

	log.Println("Setting up handlers")
	r.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling '/'")
		w.Write([]byte("Hello World"))
	}).Methods(http.MethodGet)

	log.Println("Setting up CORS")
	r.Use(mux.CORSMethodMiddleware(r))

	// Start the HTTP server on port 8090
	log.Println("Starting HTTP server on http://localhost:8090" + basePath)
	log.Fatal(http.ListenAndServe(":8090", r))
}
