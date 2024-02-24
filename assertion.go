package main

import (
	"net/http"
)

func BeginLogin(w http.ResponseWriter, r *http.Request) {
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

	options, sessionData, err := webAuthn.BeginDiscoverableLogin()
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

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
}
