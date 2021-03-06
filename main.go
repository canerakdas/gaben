package main

import (
	"github.com/canerakdas/gaben/controllers"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
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

	// User operations
	r.HandleFunc("/user",controllers.GetUser).Methods("GET")
	r.HandleFunc("/user",controllers.PostUser).Methods("POST")
	r.HandleFunc("/user",controllers.PatchUser).Methods("PATCH")
	r.HandleFunc("/user",controllers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("template/styles/"))))

	controllers.CollectHeader()
	c := cron.New()
	c.AddFunc("@every 0h10m0s", controllers.CronGameView)
	c.Start()

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
