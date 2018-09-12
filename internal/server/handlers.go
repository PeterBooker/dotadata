package server

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/PeterBooker/dota2data/internal/tmpls"
	"github.com/PeterBooker/dota2data/internal/data"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

var templates *template.Template

func init() {
	var err error
	tmpl := template.New("website")
	templates, err = vfstemplate.ParseGlob(tmpls.Assets, tmpl, "*.html")
	if err != nil {
		log.Fatalf("Failed to read template files: %s\n", err)
	}
}

// Render HTML pages
func render(w http.ResponseWriter, template string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Vary", "Accept-Encoding")

	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type errResponse struct {
	Code string `json:"code,omitempty"`
	Err  string `json:"error"`
}

type Page struct {
	Name        string
	Title       string
	Description string
	Path        string
}

type App struct {
	Name    string
	Version string
	Host    string
}

func (s *Server) index() http.HandlerFunc {
	page := Page{
		Name:        "index",
		Title:       "Dota2 Data",
		Description: "Display Dota 2 data on your website.",
		Path:        "/",
	}
	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		Host:    s.Config.Host,
	}

	data := struct {
		Page Page
		App  App
	}{
		page,
		app,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "indexPage", data)
	}
}

func (s *Server) docs() http.HandlerFunc {
	page := Page{
		Name:        "docs",
		Title:       "Docs - Dota 2 Data",
		Description: "Learn how to display Dota2 data on your website.",
		Path:        "/docs",
	}
	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		Host:    s.Config.Host,
	}

	data := struct {
		Page Page
		App  App
	}{
		page,
		app,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "docsPage", data)
	}
}

func (s *Server) notFound() http.HandlerFunc {
	page := Page{
		Name:        "notfound",
		Title:       "404 - Dota 2 Data",
		Description: "Not Found",
	}
	app := App{
		Name:    s.Config.Name,
		Version: s.Config.Version,
		Host:    s.Config.Host,
	}

	data := struct {
		Page Page
		App  App
	}{
		page,
		app,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		page.Path = r.URL.Path
		w.WriteHeader(http.StatusNotFound)
		render(w, "notFoundPage", data)
	}
}

// getHeroes fetches a list of all Heroes
func (s *Server) getHeroes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Data.GetHeroes()
		if err != nil {
			var resp errResponse
			resp.Err = fmt.Sprintf("Hero list not found: %s", err)
			w.WriteHeader(http.StatusNotFound)
			writeResp(w, resp)
			return
		}

		w.Write(data)
	}
}

// getHero fetches the data for a hero
func (s *Server) getHero() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := data.ValidateLang(r.URL.Query().Get("lang"))

		if heroName := chi.URLParam(r, "name"); heroName != "" {
			var name string
			// Check if Name is an ID instead
			if _, err := strconv.ParseInt(heroName, 10, 64); err == nil {
				name, err = s.Data.GetHeroByID(heroName)
				if err != nil {
					var resp errResponse
					resp.Err = fmt.Sprintf("Hero %s not found", heroName)
					w.WriteHeader(http.StatusNotFound)
					writeResp(w, resp)
					return
				}
			} else {
				name = data.HeroNameFromAlias(heroName)
			}
			
			data, err := s.Data.GetHero(name, lang)
			if err != nil {
				var resp errResponse
				resp.Err = fmt.Sprintf("Hero %s not found", heroName)
				w.WriteHeader(http.StatusNotFound)
				writeResp(w, resp)
				return
			}
			w.Write(data)
			return
		}

		var resp errResponse
		resp.Err = "You must specify a valid Hero name"
		w.WriteHeader(http.StatusBadRequest)
		writeResp(w, resp)
	}
}

// getAbilities fetches a list of all Abilities
func (s *Server) getAbilities() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Data.GetAbilities()
		if err != nil {
			var resp errResponse
			resp.Err = fmt.Sprintf("Ability list not found: %s", err)
			w.WriteHeader(http.StatusNotFound)
			writeResp(w, resp)
			return
		}

		w.Write(data)
	}
}

// getAbility fetches the data for an ability
func (s *Server) getAbility() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := data.ValidateLang(r.URL.Query().Get("lang"))

		if abilityName := chi.URLParam(r, "name"); abilityName != "" {
			data, err := s.Data.GetAbility(abilityName, lang)
			if err != nil {
				var resp errResponse
				resp.Err = fmt.Sprintf("Ability %s not found", abilityName)
				w.WriteHeader(http.StatusNotFound)
				writeResp(w, resp)
				return
			}
			w.Write(data)
			return
		}

		var resp errResponse
		resp.Err = "You must specify a valid Ability name"
		w.WriteHeader(http.StatusBadRequest)
		writeResp(w, resp)
	}
}

// getItems fetches a list of all Items
func (s *Server) getItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Data.GetItems()
		if err != nil {
			var resp errResponse
			resp.Err = fmt.Sprintf("Item list not found: %s", err)
			w.WriteHeader(http.StatusNotFound)
			writeResp(w, resp)
			return
		}

		w.Write(data)
	}
}

// getItem fetches the data for an item
func (s *Server) getItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := data.ValidateLang(r.URL.Query().Get("lang"))

		if itemName := chi.URLParam(r, "name"); itemName != "" {
			name := data.ItemNameFromAlias(itemName)
			data, err := s.Data.GetItem(name, lang)
			if err != nil {
				var resp errResponse
				resp.Err = fmt.Sprintf("Item %s not found", name)
				w.WriteHeader(http.StatusNotFound)
				writeResp(w, resp)
				return
			}
			w.Write(data)
			return
		}

		var resp errResponse
		resp.Err = "You must specify a valid Item name"
		w.WriteHeader(http.StatusBadRequest)
		writeResp(w, resp)
	}
}

// getUnits fetches a list of all Units
func (s *Server) getUnits() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Data.GetUnits()
		if err != nil {
			var resp errResponse
			resp.Err = fmt.Sprintf("Unit list not found: %s", err)
			w.WriteHeader(http.StatusNotFound)
			writeResp(w, resp)
			return
		}

		w.Write(data)
	}
}

// getUnit fetches the data for an item
func (s *Server) getUnit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := data.ValidateLang(r.URL.Query().Get("lang"))

		if unitName := chi.URLParam(r, "name"); unitName != "" {
			data, err := s.Data.GetUnit(unitName, lang)
			if err != nil {
				var resp errResponse
				resp.Err = fmt.Sprintf("Unit %s not found", unitName)
				w.WriteHeader(http.StatusNotFound)
				writeResp(w, resp)
				return
			}
			w.Write(data)
			return
		}

		var resp errResponse
		resp.Err = "You must specify a valid Unit name"
		w.WriteHeader(http.StatusBadRequest)
		writeResp(w, resp)
	}
}
