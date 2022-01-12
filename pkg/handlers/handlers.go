package handlers

import (
	"net/http"

	"github.com/tunedev/bookings_in_go/pkg/config"
	"github.com/tunedev/bookings_in_go/pkg/models"
	"github.com/tunedev/bookings_in_go/pkg/renders"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	renders.RenderPageTemplate(w, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringData := make(map[string]string)
	stringData["test"] = "Som random test to learn about passing data to a template"

	stringData["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	
	renders.RenderPageTemplate(w, "about", &models.TemplateData{
		StringMap: stringData,
	})
}