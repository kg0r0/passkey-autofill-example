package main

import (
	"fmt"
	"sync"
)

type userdb struct {
	users map[string]*User
	mu    sync.RWMutex
}

var usersDB *userdb = &userdb{
	users: make(map[string]*User),
}

func (db *userdb) GetUser(id string) (*User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (db *userdb) AddUser(user *User) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, ok := db.users[string(user.WebAuthnID())]
	if ok {
		return fmt.Errorf("user already exists")
	}
	_, ok = db.users[string(user.username)]
	if ok {
		return fmt.Errorf("user already exists")
	}
	db.users[string(user.WebAuthnID())] = user
	db.users[user.username] = user
	return nil
}
