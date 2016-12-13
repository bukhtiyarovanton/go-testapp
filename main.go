// main
package main

import (
	//	"encoding/json"
	"log"
	"net/http"
	"os"
	//"time"
	//	"fmt"

	//"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/handlers"
	//	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := NewRouter()

	//log.Fatal(http.ListenAndServe(":3000", handlers.CombinedLoggingHandler(os.Stdout, router)))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
