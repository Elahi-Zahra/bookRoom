package handlers

import (
	"github.com/Elahi-Zahra/bookRoom/pkg/config"
	"github.com/Elahi-Zahra/bookRoom/pkg/models"
	"github.com/Elahi-Zahra/bookRoom/pkg/render"
	"net/http"
)



//Repo the repository used by handler
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo create new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandler set the repository from handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	render.RenderTemplate(w, "home.page.tmpl",&models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap :=make(map[string]string)
	stringMap["test"]="hello world"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIP

	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
	//render.RenderTemplate(w, "about.page.tmpl")
}
