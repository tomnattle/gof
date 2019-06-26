package main

import (
	"github.com/tomnattle/gof"
	"github.com/tomnattle/gof/cms/controller/article"
)

func main() {
	app := gof.New()
	app.Rest("/article", &article.Article{})
	app.Start()
}
