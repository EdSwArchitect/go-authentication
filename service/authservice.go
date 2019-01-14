package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	authentication "github.com/EdSwArchitect/go-authentication/db"

	"github.com/gorilla/mux"
)

// Return the root path
func getRootPath(w http.ResponseWriter, r *http.Request) {
	contents := `{ "response" : "Hi, Ed! Webbie Root path"}`
	json.NewEncoder(w).Encode(contents)
}

// Return the root path
func getUser(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	fmt.Printf("Params: %v\n", params)

	fmt.Printf("val is: '%s'\n", params["name"])

	contents := fmt.Sprintf("{ \"user\" : \"%s\"", params["name"])

	log.Printf("Looking for user '%s'", contents)

	name := params["name"]

	user, err := authentication.GetUserByName(name)

	if err != nil {
		log.Printf("User '%s' not found", name)
	}

	json.NewEncoder(w).Encode(user)
}

// Server sets up the RESTful server with the URIs given
func Server(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/", getRootPath).Methods("GET")
	router.HandleFunc("/auth/user/{name}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
