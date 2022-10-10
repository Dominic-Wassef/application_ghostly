package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/dominic-wassef/ghostly"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	gho := &ghostly.Ghostly{}
	err = gho.New(path)
	if err != nil {
		log.Fatal(err)
	}

	gho.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: gho,
	}

	myHandlers := &handlers.Handlers{
		App: gho,
	}

	app := &application{
		App:        gho,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
