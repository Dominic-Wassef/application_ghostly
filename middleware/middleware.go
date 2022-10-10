package middleware

import (
	"myapp/data"

	"github.com/dominic-wassef/ghostly"
)

type Middleware struct {
	App    *ghostly.Ghostly
	Models data.Models
}
