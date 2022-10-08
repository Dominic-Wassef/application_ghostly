package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/dominic-wassef/ghostly"
)

type application struct {
	App *ghostly.Ghostly
	Handlers *handlers.Handlers
	Models data.Models
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
