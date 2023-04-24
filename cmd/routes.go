package cmd

import (
	"net/http"

	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// route for user
	mux.Post("/register", controllers.RepoUser.Register)
	mux.Get("/login", controllers.RepoUser.Login)
	mux.Get("/users", controllers.RepoUser.GetUser)

	// route for store
	mux.Post("/stores", controllers.RepoStore.SetStore)
	mux.Get("/stores", controllers.RepoStore.GetStore)
	mux.Get("/stores/id", controllers.RepoStore.GetOneStore)

	// route for status
	mux.Post("/status", controllers.RepoStatus.SetStatus)
	mux.Get("/status", controllers.RepoStatus.GetStatus)
	mux.Get("/status/id", controllers.RepoStatus.GetOneStatus)

	// route for goods
	mux.Post("/goods", controllers.RepoGoods.SetGoods)
	mux.Get("/goods", controllers.RepoGoods.GetGoods)
	mux.Get("/goods/id", controllers.RepoGoods.GetOneGoods)

	return mux
}
