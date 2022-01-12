package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tunedev/bookings_in_go/pkg/config"
	"github.com/tunedev/bookings_in_go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(SaySomethingInTheConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}