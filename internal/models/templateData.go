package models

import "github.com/tunedev/bookings_in_go/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	Intmap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
