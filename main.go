package main

import (
	"github.com/canerakdas/game_crawl/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/games/{id}", controllers.IGetGame).Methods("GET")
	r.HandleFunc("/games/{id}/{name}", controllers.IGetGame).Methods("GET")
	r.HandleFunc("/games", controllers.GetGame).Methods("GET")

	//r.HandleFunc("/catalog/",controllers.GetCatalog).Methods("GET")
	r.HandleFunc("/catalog/{id}/{name}", controllers.GetCatalog).Methods("GET")

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("template/styles/"))))

	controllers.CollectHeader()
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
