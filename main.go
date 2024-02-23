package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	webAuthn *webauthn.WebAuthn
	err      error
)

type FIDO2Response struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func jsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func main() {
	wconfig := &webauthn.Config{
		RPDisplayName: "Go WebAuthn",
		RPID:          "localhost",
		RPOrigin:      "http://localhost:8080",
	}
	if webAuthn, err = webauthn.New(wconfig); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/attestation/options", BeginRegistration)
	http.HandleFunc("/attestation/result", FinishRegistration)

	http.HandleFunc("/assertion/options", BeginRegistration)
	http.HandleFunc("/assertion/result", FinishRegistration)

	http.Handle("/", http.FileServer(http.Dir("./templates")))

	serverAddr := "localhost:8080"

	log.Println("Listening on http://" + serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
