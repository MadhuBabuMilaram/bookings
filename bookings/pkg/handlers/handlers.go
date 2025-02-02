package handlers

import (
	"bookings/pkg/config"
	modelsTemplate "bookings/pkg/models"
	"bookings/pkg/renders"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

var Repo *Repository

func NewHandler(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	stringMap := map[string]string{}
	stringMap["test"] = "Hello again from home data"
	renders.RenderTemplate(w, "homepage.html", &modelsTemplate.TemplateData{StringMap: stringMap})

}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["test"] = "Hello again from about data "
	stringMap["remoteIP"] = remoteIP
	renders.RenderTemplate(w, "about.html", &modelsTemplate.TemplateData{StringMap: stringMap})
}
