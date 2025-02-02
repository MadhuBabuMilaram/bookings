package renders

import (
	"bookings/pkg/config"
	modelsTemplate "bookings/pkg/models"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, templ string) {
	parsedFiles, _ := template.ParseFiles("./templates/"+templ, "./templates/baselayout.html")
	err := parsedFiles.Execute(w, nil)
	if err != nil {
		fmt.Println("failed to execture", err)
	}
}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *modelsTemplate.TemplateData) *modelsTemplate.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *modelsTemplate.TemplateData) {
	//create TemplateCache
	// get the templateCahce from appconfig
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	// render template.
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	log.Println("create template cache")
	// get all filenames from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	//range through the pages

	for _, page := range pages {
		//log.Println(page)
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
