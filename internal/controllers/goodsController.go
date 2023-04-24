package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/models"
)

// RepoGoods is repository used by user controller
var RepoGoods *RepositoryGoods

// RepositoryGoods is repository type
type RepositoryGoods struct {
	App *config.AppConfig
}

// NewGoodsRepo create new repository
func NewGoodsRepo(a *config.AppConfig) *RepositoryGoods {
	return &RepositoryGoods{
		App: a,
	}
}

// NewGoodsController create new users
func NewGoodsController(r *RepositoryGoods) {
	RepoGoods = r
}

// SetGoods is function to create store
func (m *RepositoryGoods) SetGoods(w http.ResponseWriter, r *http.Request) {
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

// GetGoods is function to get all store
func (m *RepositoryGoods) GetGoods(w http.ResponseWriter, r *http.Request) {
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

// GetOneGoods is function to get one store
func (m *RepositoryGoods) GetOneGoods(w http.ResponseWriter, r *http.Request) {
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
