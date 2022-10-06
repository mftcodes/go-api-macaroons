package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/macaroon.v2"
)

type localMacaroon struct {
	Secret   string `json:"secret"`
	Id       string `json:"id"`
	Location string `json:"location"`
}

type caveat struct {
	Id             []byte
	VerificationId []byte
	Location       string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func newMac(key []byte, id []byte, location string) *macaroon.Macaroon {
	m, err := macaroon.New(key, id, location, macaroon.LatestVersion)
	if err != nil {
		printError(err)
	}
	return m
}

func doMacaroon(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var mac localMacaroon
	json.Unmarshal(reqBody, &mac)

	m := newMac([]byte(mac.Secret), []byte(mac.Id), mac.Location)

	fpc := caveat{
		Id:             []byte("Has_Parent_Org"),
		VerificationId: []byte(""),
		Location:       "www.test.com/two",
	}

	m.AddFirstPartyCaveat(fpc.Id)

	json.NewEncoder(w).Encode(m)
}

func printError(err error) {
	fmt.Println(err)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/macaroon", doMacaroon).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
