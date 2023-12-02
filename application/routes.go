package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"web-calculator/handler"
	"web-calculator/logger"
)

// load routes handler
func (a *App) loadRoutes() {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger)

	// create index route
	router.Route("/", a.equationRoute)

	// set router
	a.router = router

}

func (a *App) equationRoute(router chi.Router) {

	// create handler function
	eqHandler := &handler.EquationHandler{
		Logger: &logger.Logger{
			Client: a.db,
		},
	}

	// add handler to router
	router.Post("/", eqHandler.ProcessEquation)
}