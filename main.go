package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/dominic-wassef/ghostly"
)

type application struct {
	App        *ghostly.Ghostly
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
