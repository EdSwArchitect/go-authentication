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

// Return the root path
func addUser(w http.ResponseWriter, r *http.Request) {

	// a map of values
	params := mux.Vars(r)

	fmt.Printf("Params: %+v\n", params)

	form := r.Form

	fmt.Printf("The form stuff: %+v\n", form)

	// body, err := ioutil.ReadAll(r.Body)

	// if err != nil {
	// 	log.Fatal("Unable to read the body: ", err)
	// }

	// log.Printf("The body is: %s\n", body)

	var theUser authentication.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&theUser)

	if err != nil {
		log.Fatal("Decode error: ", err)
	}

	fmt.Printf("the decoded value: %+v\n", theUser)

	// fmt.Printf("val is: '%s'\n", params["name"])

	// contents := fmt.Sprintf("{ \"user\" : \"%s\"", params["name"])

	// log.Printf("Looking for user '%s'", contents)

	// name := params["name"]

	// user, err := authentication.GetUserByName(name)

	// if err != nil {
	// 	log.Printf("User '%s' not found", name)
	// }

	json.NewEncoder(w).Encode(`{"results" : "good"}`)
}

// Server sets up the RESTful server with the URIs given
func Server(port int) {
	router := mux.NewRouter()

	router.HandleFunc("/", getRootPath).Methods("GET")
	router.HandleFunc("/auth/user/{name}", getUser).Methods("GET")
	router.HandleFunc("/auth/user/{name}", addUser).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
