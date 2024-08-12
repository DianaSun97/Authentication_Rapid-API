package handlers

import (
	"github.com/DianaSun97/elluliin_booking/pkg/config"
	"github.com/DianaSun97/elluliin_booking/pkg/forms"
	"github.com/DianaSun97/elluliin_booking/pkg/helpers"
	"github.com/DianaSun97/elluliin_booking/pkg/models"
	"github.com/DianaSun97/elluliin_booking/render"
	"net/http"
)

// Repo
var Repo *Repository

// Repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.gohtml", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.gohtml", &models.TemplateData{})
}

// Generals
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals-page.gohtml", &models.TemplateData{})
}

// Majors
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors-page.gohtml", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{
		Form: nil,
		Data: data,
	})
}

// CreateReservation handles the reservation creation
func (m *Repository) CreateReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

}
