package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func SaySomethingInTheConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Testing Middleware")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csfrHandler := nosurf.New(next)

	csfrHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csfrHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}