package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	graphiccards "github.com/sigsant/graphic_card_catalog_app/v2/api/graphicCards"
)

func cardCatalog( w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/graphic-cards" {
		http.Error(w, "404: Page not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

	json.NewEncoder(w).Encode(graphiccards.Card)
}

func HandlerRequest(mreq *mux.Router){
	mreq.HandleFunc("/graphic-cards", cardCatalog)
}

func main(){

	muxRequest := mux.NewRouter().StrictSlash(true)
	HandlerRequest(muxRequest)
	fmt.Println("Se ha iniciado el servidor en el puerto http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", muxRequest))
}

func CheckerError(err error){
	if err != nil {
		log.Fatal(err)
	}
}