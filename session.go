package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/go-webauthn/webauthn/webauthn"
)

type sessiondb struct {
	sessions map[string]*webauthn.SessionData
	mu       sync.RWMutex
}

var sessionDb *sessiondb = &sessiondb{
	sessions: make(map[string]*webauthn.SessionData),
}

func (db *sessiondb) GetSession(sessionID string) (*webauthn.SessionData, error) {

	db.mu.Lock()
	defer db.mu.Unlock()

	session, ok := db.sessions[sessionID]
	if !ok {
		return nil, fmt.Errorf("error getting session '%s': does not exist", sessionID)
	}
	return session, nil
}

func (db *sessiondb) DeleteSession(sessionID string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.sessions, sessionID)
}

func (db *sessiondb) StartSession(data *webauthn.SessionData) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	id, _ := random(32)
	db.sessions[id] = data
	return id
}

func random(length int) (string, error) {
	randomData := make([]byte, length)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(randomData), nil
}
