package app

import (
	"net/http"

	"github.com/brentchesny/go-inertia/inertia"
	"github.com/gorilla/mux"
)

type App struct {
	mux *mux.Router
}

func New() *App {
	return &App{
		mux: mux.NewRouter(),
	}
}

func (a *App) Run() {
	a.registerRoutes()
	http.ListenAndServe(":8080", a.mux)
}

func (a *App) registerRoutes() {
	a.mux.HandleFunc("/", HandleHome)
	a.mux.HandleFunc("/second", HandleSecond)
	a.mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	var name string
	if names := r.URL.Query()["name"]; len(names) > 0 {
		name = names[0]
	}

	inertia.Render(w, r, "Home", map[string]interface{}{
		"name": name,
	})
}

func HandleSecond(w http.ResponseWriter, r *http.Request) {
	inertia.Render(w, r, "Second", nil)
}
