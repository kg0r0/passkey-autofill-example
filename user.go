package main

import "github.com/go-webauthn/webauthn/webauthn"

type User struct {
	id          uint64
	username    string
	displayName string
	credentials []webauthn.Credential
}
