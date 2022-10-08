package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
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

	myHandlers := &handlers.Handlers{
		App: gho,
	}

	app := &application{
		App:      gho,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models

	return app
}
