package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/models"
)

// Repo is repository used by user controller
var Repo *Repository

// Repository is repository type
type Repository struct {
	App *config.AppConfig
}

// NewUserRepo create new repository
func NewUserRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewUserController create new users
func NewUserController(r *Repository) {
	Repo = r
}

func (m *Repository) Signup(w http.ResponseWriter, r *http.Request) {
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
