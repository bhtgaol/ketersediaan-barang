package controllers

import (
	"encoding/json"
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
	resp := models.User{
		Name:         "Boy",
		PhoneNumber:  "08123123123",
		EmailAddress: "boy@gmail.com",
		Password:     "test",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Login is function to login user
func (m *RepositoryUser) Login(w http.ResponseWriter, r *http.Request) {
	resp := models.User{
		Name:         "Boy",
		PhoneNumber:  "08123123123",
		EmailAddress: "boy@gmail.com",
		Password:     "test",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// GetUser is function to Get all user
func (m *RepositoryUser) GetUser(w http.ResponseWriter, r *http.Request) {
	resp := models.User{
		Name:         "Boy",
		PhoneNumber:  "08123123123",
		EmailAddress: "boy@gmail.com",
		Password:     "test",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
