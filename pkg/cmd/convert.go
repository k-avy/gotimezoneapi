package cmd

import (
	"fmt"
	"log"
	"net/http"

	"example/github.com/k-avy/gotimezoneapi/pkg/convtimezone"
	"example/github.com/k-avy/gotimezoneapi/pkg/convzone"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func ConvertZone() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/convert", convtimezone.Convert).Methods("GET")
	router.HandleFunc("/converttz", convzone.ConvertTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
