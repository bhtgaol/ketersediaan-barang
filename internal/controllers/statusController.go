package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/models"
)

// RepoStatus is repository used by user controller
var RepoStatus *RepositoryStatus

// RepositoryStatus is repository type
type RepositoryStatus struct {
	App *config.AppConfig
}

// NewStatusRepo create new repository
func NewStatusRepo(a *config.AppConfig) *RepositoryStatus {
	return &RepositoryStatus{
		App: a,
	}
}

// NewStatusController create new users
func NewStatusController(r *RepositoryStatus) {
	RepoStatus = r
}

// SetStatus is function to create store
func (m *RepositoryStatus) SetStatus(w http.ResponseWriter, r *http.Request) {
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

// GetStatus is function to get all store
func (m *RepositoryStatus) GetStatus(w http.ResponseWriter, r *http.Request) {
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

// GetOneStatus is function to get one store
func (m *RepositoryStatus) GetOneStatus(w http.ResponseWriter, r *http.Request) {
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
