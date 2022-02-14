package handlers

import (
	"encoding/json"
	"github.com/tunedev/bookings_in_go/internal/forms"
	"log"
	"net/http"

	"github.com/tunedev/bookings_in_go/internal/config"
	"github.com/tunedev/bookings_in_go/internal/models"
	"github.com/tunedev/bookings_in_go/internal/renders"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
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

type generalAvailabilityRes struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	resp := generalAvailabilityRes{
		Start: r.PostFormValue("start"),
		End:   r.PostFormValue("end"),
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println("Error parsing general availability", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

type singleAvailabilityRes struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

func (m *Repository) PostAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := singleAvailabilityRes{
		OK:      true,
		Message: "Something to test",
		Start:   r.Form.Get("startModal"),
		End:     r.PostFormValue("endModal"),
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println("An error occured while parson json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringData := make(map[string]string)
	stringData["test"] = "Som random test to learn about passing data to a template"

	stringData["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	renders.RenderPageTemplate(w, r, "about", &models.TemplateData{
		StringMap: stringData,
	})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	renders.RenderPageTemplate(w, r, "make-reservation", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
}
