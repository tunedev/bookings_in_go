package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tunedev/bookings_in_go/pkg/config"
	"github.com/tunedev/bookings_in_go/pkg/handlers"
	"github.com/tunedev/bookings_in_go/pkg/renders"
)

const portNumber = ":3000"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	ts, err := renders.CreateTemplateSet()
	if err != nil {
		fmt.Println("Error occured while creating templateSet")
	}

	app.TemplateCache = ts
	app.UseConfig = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	fmt.Printf("My app server running on port http://localhost%s \n", portNumber)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}
