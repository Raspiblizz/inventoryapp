package main 

import (
	 		"fmt"
			"log"
			"net/http"
			"github.com/gorilla/mux"
)

func handleRequests() {
		router := mux.NewRouter().StrictSlash(true)

		router.HandleFunc("/", homePage).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))

}


func homePage(w http.ResponsWriter, r *http.Request) {
       fmt.Fprintf(w, "Endpoint called: homePage()")
  
}

func main() {
    handleRequests()
		
}