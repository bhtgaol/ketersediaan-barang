package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/models"
)

// RepoStore is repository used by user controller
var RepoStore *RepositoryStore

// RepositorySotre is repository type
type RepositoryStore struct {
	App *config.AppConfig
}

// NewUserRepo create new repository
func NewStoreRepo(a *config.AppConfig) *RepositoryStore {
	return &RepositoryStore{
		App: a,
	}
}

// NewUserController create new users
func NewStoreController(r *RepositoryStore) {
	RepoStore = r
}

// SetStore is function to create store
func (m *RepositoryStore) SetStore(w http.ResponseWriter, r *http.Request) {
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

// GetStore is function to get all store
func (m *RepositoryStore) GetStore(w http.ResponseWriter, r *http.Request) {
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

// GetOneStore is function to get one store
func (m *RepositoryStore) GetOneStore(w http.ResponseWriter, r *http.Request) {
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
