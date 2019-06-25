package main

import (
	"github.com/tomnattle/gof"
	"github.com/tomnattle/gof/cms/controller/article"
)

func main() {
	app := gof.New()
	app.Get("/article", article.List)
	app.Post("/article", article.List)
	app.Start()
}
