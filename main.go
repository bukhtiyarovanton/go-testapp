// main
package main

import (
	//	"encoding/json"
	"log"
	"net/http"
	//"os"
	//"time"
	//	"fmt"

	//"github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/handlers"
	//	"github.com/gorilla/mux"
)

func main() {

	router := NewRouter()

	//log.Fatal(http.ListenAndServe(":3000", handlers.CombinedLoggingHandler(os.Stdout, router)))
	log.Fatal(http.ListenAndServe(":3000", router))
}
