package main

import (
	"log"
	
	"github.com/gorilla/mux"

	"github.com/TursunovImran/mini-projects_go/web_wordfile_search/pkg"
	webwordfilesearch "github.com/TursunovImran/mini-projects_go/web_wordfile_search"
)

func main() {

	router := mux.NewRouter()
    router.HandleFunc("/files/search", pkg.SearchWordFile).Methods("GET")

	srv := new(webwordfilesearch.Server)
	if err := srv.Run("8080", router); err != nil {
		log.Fatalf("error ocured while running http server %s", err.Error())
	}
}