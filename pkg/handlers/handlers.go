package handlers

import (
	"fmt"
	"log"
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
	renders.RenderPageTemplate(w, r, "home", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	renders.RenderPageTemplate(w, r, "contact", &models.TemplateData{})
}
func (m *Repository) Captain(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	renders.RenderPageTemplate(w, r, "captain", &models.TemplateData{})
}
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	renders.RenderPageTemplate(w, r, "major", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renders.RenderPageTemplate(w, r, "availability", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method)
	log.Println("It is getting here", r.Form.Get("start"))
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("The start Date: %s, and End Date: %s", start, end)))
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringData := make(map[string]string)
	stringData["test"] = "Som random test to learn about passing data to a template"

	stringData["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	
	renders.RenderPageTemplate(w, r, "about", &models.TemplateData{
		StringMap: stringData,
	})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request){
	
	renders.RenderPageTemplate(w, r, "make-reservation", &models.TemplateData{})
}