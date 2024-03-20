package main

import (
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

func AssertionOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "Invalid request method",
		}, http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Some options are complemented in the frontend
	// Ref: https://github.com/MasterKale/SimpleWebAuthn/blob/5229cebbcc2d087b7eaaaeb9886f53c9e1d93522/packages/browser/src/methods/startAuthentication.ts#L72-L76
	options, sessionData, err := webAuthn.BeginDiscoverableLogin()
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "authentication",
		Value: sessionDb.StartSession(sessionData),
		Path:  "/",
	})
	jsonResponse(w, options, http.StatusOK)
}

func AssertionResult(w http.ResponseWriter, r *http.Request) {
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

	cookie, err := r.Cookie("authentication")
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

	credential, err := webAuthn.FinishDiscoverableLogin(func(rawId []byte, userhandle []byte) (user webauthn.User, err error) {
		return usersDB.GetUser(string(userhandle))
	}, *sessionData, r)
	if err != nil {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if !credential.Flags.UserVerified {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "user was not verified",
		}, http.StatusBadRequest)
		return
	}

	if credential.Authenticator.CloneWarning {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "authenticator is cloned",
		}, http.StatusBadRequest)
		return
	}

	sessionDb.DeleteSession(cookie.Value)

	jsonResponse(w, FIDO2Response{
		Status:       "ok",
		ErrorMessage: "",
	}, http.StatusOK)
}
