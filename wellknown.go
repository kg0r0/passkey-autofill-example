package main

import (
	"log/slog"
	"net/http"
)

type WellknownWebAuthnResponse struct {
	Origins []string `json:"origins"`
}

// Ref: https://passkeys.dev/docs/advanced/related-origins/#relying-party-changes
func WebAuthn(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Host + "/.well-known/webauthn")
	if r.Method != http.MethodGet {
		jsonResponse(w, FIDO2Response{
			Status:       "failed",
			ErrorMessage: "Invalid request method",
		}, http.StatusBadRequest)
		return
	}
	jsonResponse(w, WellknownWebAuthnResponse{
		Origins: []string{
			"http://localhost:8080",
			"https://shopping.com",
			"https://myshoppingrewards.com",
			"https://myshoppingcreditcard.com",
			"https://myshoppingtravel.com",
			"https://shopping.co.uk",
			"https://shopping.co.jp",
			"https://shopping.ie",
			"https://shopping.ca",
		},
	}, http.StatusOK)
}
