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

var (
	rpid = "shopping.co.uk"
)

type Params struct {
	Username string
}

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
		RPID:          rpid,
		RPOrigins:     []string{"https://shopping.co.uk", "https://shopping.com"},
	}
	if webAuthn, err = webauthn.New(wconfig); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/attestation/options", AttestationOptions)
	http.HandleFunc("/attestation/result", AttestationResult)

	http.HandleFunc("/assertion/options", AssertionOptions)
	http.HandleFunc("/assertion/result", AssertionResult)

	http.HandleFunc("/.well-known/webauthn", WebAuthn)

	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	})

	port := "443"

	log.Fatal(http.ListenAndServeTLS(":"+port, "./certs/shopping.com+1.pem", "./certs/shopping.com+1-key.pem", nil))
}
