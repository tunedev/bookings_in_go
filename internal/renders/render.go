package renders

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/tunedev/bookings_in_go/internal/config"
	"github.com/tunedev/bookings_in_go/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderPageTemplate(w http.ResponseWriter, r *http.Request, page string, td *models.TemplateData) {
	var ts map[string]*template.Template
	if app.UseCache {
		ts = app.TemplateCache
	}else {
		ts, _ = CreateTemplateSet()
	}

	templatePath :=  page+ ".page.tmpl"
	t, ok := ts[templatePath]
	if !ok {
		log.Fatal(page+": page does not exist")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Printf("An error occured while writing to response writer")
	}
}

func CreateTemplateSet() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages{
		pageName := filepath.Base(page)

		ts, err := template.New(pageName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[pageName] = ts
	}
	return myCache, nil
}
