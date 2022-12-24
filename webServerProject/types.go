package main

import (
	"encoding/json"
	"net/http"
)

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// User type
type User struct {
	Name  string `json:"name"` // tag
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

// MetaData type
type MetaData interface{}
