package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
)

func AttestationOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	var p Params

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	username := p.Username
	if username == "" {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "Missing username",
		}, http.StatusBadRequest)
		return
	}

	user, err := usersDB.GetUser(username)
	if err != nil {
		displayName := strings.Split(username, "@")[0]
		user = NewUser(username, displayName)
		usersDB.AddUser(user)
	}
	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = user.CredentialExcludeList()
	}
	options, sessionData, err := webAuthn.BeginRegistration(user, registerOptions)
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "registration",
		Value: sessionDb.StartSession(sessionData),
		Path:  "/",
	})
	jsonResponse(w, options, http.StatusOK)
}

func AttestationResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	cookie, err := r.Cookie("registration")
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	sessionData, err := sessionDb.GetSession(cookie.Value)
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	user, err := usersDB.GetUser(string(sessionData.UserID))
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	credential, err := webAuthn.FinishRegistration(user, *sessionData, r)
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if !credential.Flags.UserPresent || !credential.Flags.UserVerified {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "user was not verified",
		}, http.StatusBadRequest)
		return
	}

	user.AddCredential(*credential)

	sessionDb.DeleteSession(cookie.Value)

	jsonResponse(w, FIDO2Response{
		Status:       "ok",
		ErrorMessage: "",
	}, http.StatusOK)
}
