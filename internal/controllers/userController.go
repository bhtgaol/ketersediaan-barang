package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/models"
)

// RepoUser is repository used by user controller
var RepoUser *RepositoryUser

// RepositoryUser is repository type
type RepositoryUser struct {
	App *config.AppConfig
}

// NewUserRepo create new repository
func NewUserRepo(a *config.AppConfig) *RepositoryUser {
	return &RepositoryUser{
		App: a,
	}
}

// NewUserController create new users
func NewUserController(r *RepositoryUser) {
	RepoUser = r
}

// Register is function use to register user
func (m *RepositoryUser) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello World!")
}

// Login is function to login user
func (m *RepositoryUser) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	staticUser := models.User{
		Name:         "Boy",
		PhoneNumber:  "08123123123",
		EmailAddress: "boy@gmail.com",
		Password:     "test",
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// checking email
	if staticUser.EmailAddress != user.EmailAddress {
		http.Error(w, "Email yang anda masukkan tidak ditemmukan", http.StatusBadRequest)
		return
	}

	// checking password
	if staticUser.Password != user.Password {
		http.Error(w, "Password yang anda masukkan salah", http.StatusBadRequest)
		return
	}
}

// GetUser is function to Get all user
func (m *RepositoryUser) GetUser(w http.ResponseWriter, r *http.Request) {
	var respslice []models.User
	resp := models.User{
		Name:         "Boy",
		PhoneNumber:  "08123123123",
		EmailAddress: "boy@gmail.com",
		Password:     "test",
	}

	respslice = append(respslice, resp)
	respslice = append(respslice, resp)
	respslice = append(respslice, resp)

	out, err := json.MarshalIndent(respslice, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
