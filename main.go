package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

var (
	webAuthn *webauthn.WebAuthn
	err      error
)

func main() {
	wconfig := &webauthn.Config{
		RPDisplayName: "Go WebAuthn",
		RPID:          "localhost",
		RPOrigin:      "http://localhost:8080",
	}
	if webAuthn, err = webauthn.New(wconfig); err != nil {
		fmt.Println(err)
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
