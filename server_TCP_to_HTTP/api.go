package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (a *api) getUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var foundUser *User
	for _, user := range users {
		if user.ID == id {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "Khong tim thay người dùng", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(foundUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *api) createUsershandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := GenerateID()

	u := User{
		ID:        id,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if err := insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("First name is required")
	}
	if u.LastName == "" {
		return errors.New("Last name is required")
	}

	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}
	users = append(users, u)
	return nil
}

func GenerateID() string {
	return uuid.New().String()
}
