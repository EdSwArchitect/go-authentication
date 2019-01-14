package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	authentication "github.com/EdSwArchitect/go-authentication/db"
	service "github.com/EdSwArchitect/go-authentication/service"
)

func main() {
	log.Printf("Hi, Ed")

	user := authentication.GetUserByID(1)

	log.Println("Object is: ", user)

	user.Disabled = false
	user.Fullname = "John Q. Public"

	first := sha256.New()
	first.Write([]byte(user.Fullname))

	something := first.Sum(nil)

	user.Hash = hex.EncodeToString(something)

	user.Salt = user.Username
	user.Disabled = false
	user.Username = fmt.Sprintf("%s-%d", user.Username, time.Now().Unix())

	val, _ := authentication.InsertUser(user)

	log.Println("The insert returned: ", val)

	log.Println("*** Starting server")

	service.Server(9090)

	log.Println("*** server started")
}
