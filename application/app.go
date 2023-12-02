package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"web-calculator/model"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	db *gorm.DB
}

// constructor
func New() *App {

	// determine db file
	var database *gorm.DB
	if os.Getenv("APP_STATE") == "docker" {
		db_file := "/var/lib/web-calculator/calculator.sqlite"

		// initialise db connection
		db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
		if err != nil {
			err_msg := fmt.Sprintf("failed to connect database: %s", err.Error())
			panic(err_msg)
		}
		database = db
	} else if os.Getenv("APP_STATE") == "cluster" {

		// postgres vars
		host := "postgres-db-postgresql.default.svc.cluster.local"
		dbname := "postgres"
		port := "5432"
		user := "postgres"
		password := os.Getenv("POSTGRES_PASSWORD")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// initialise db connection
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			err_msg := fmt.Sprintf("failed to connect database: %s", err.Error())
			panic(err_msg)
		}
		database = db

	} else {
		db_file := "calculator.sqlite"

		// initialise db connection
		db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
		if err != nil {
			err_msg := fmt.Sprintf("failed to connect database: %s", err.Error())
			panic(err_msg)
		}
		database = db
	}

	// migrate db
	database.AutoMigrate(&model.EquationModel{})

	// construct app
	app := &App{
		db: database,
	}

	// load routes
	app.loadRoutes()

	return app
}

// Run application
func (a *App) Run(ctx context.Context) error {

	// create server
	server := &http.Server{
		Addr: ":8082",
		Handler: a.router,
	}

	// create channel to listen for errors from server
	ch := make(chan error, 1)

	// start server in goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("server error: %w", err)
		}
		close(ch)
	}()

	// listen for context done or error
	select {
	case <-ctx.Done():
		// shutdown server
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)

	case err := <-ch:
		return err
	}

}