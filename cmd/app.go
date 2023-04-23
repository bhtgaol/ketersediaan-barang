package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bhtgaol/ketersediaan-barang/internal/config"
	"github.com/bhtgaol/ketersediaan-barang/internal/controllers"
	"github.com/joho/godotenv"
)

var app config.AppConfig
var session *scs.SessionManager

func AppStart() {
	portNumb := godotenv.Load(".env")

	if portNumb != nil {
		log.Fatalf("Error loading .env file")
	}

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.UseCache = false

	repo := controllers.NewUserRepo(&app)
	controllers.NewUserController(repo)

	srv := &http.Server{
		Addr:    os.Getenv("PORT_NUMBER"),
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)

}
