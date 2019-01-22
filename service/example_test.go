package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	authentication "github.com/EdSwArchitect/go-authentication/db"
)

func TestAbc(t *testing.T) {
	fmt.Println("Hi, Ed, this is a test")
	t.Log("This is something good\n")

	url := "http://localhost:9090/auth/user/gumby"

	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)

	var edwin authentication.User

	edwin.ID = 55
	edwin.Fullname = "Edwin Kenton Brown"
	edwin.Disabled = false
	edwin.Invalid = false
	edwin.Username = "edkbrown"
	edwin.Salt = "YoMomma!"

	bytes, err := json.Marshal(edwin)

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bytes)))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
