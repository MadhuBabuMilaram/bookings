package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/renders"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var portNum = ":8080"
var app config.AppConfig

func main() {

	app.Session = scs.New()
	app.Session.Cookie.Secure = false
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Persist = app.InProduction
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(("cannot create template cache"))
	}
	app.TemplateCache = tc
	app.UseCache = false
	renders.NewTemplate(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(portNum, nil)
	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
