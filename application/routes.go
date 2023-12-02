package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"web-calculator/handler"
)

// load routes handler
func (a *App) loadRoutes() {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger)

	// create index route
	router.Route("/", a.indexRoute)

	// create equation route
	router.Route("/equation", a.equationRoute)

	// set router
	a.router = router

}

func (a *App) indexRoute(router chi.Router) {

	// create handler function
	getHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}

	// add handler to router
	router.Get("/", getHandler)
}

func (a *App) equationRoute(router chi.Router) {

	// create handler function
	eqHandler := &handler.EquationHandler{}

	// add handler to router
	router.Post("/", eqHandler.ProcessEquation)
}