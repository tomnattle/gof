package main

import (
	"github.com/tomnattle/gof"
	"github.com/tomnattle/gof/cms/controller/article"
	"github.com/tomnattle/gof/module/log"
)

func main() {
	app := gof.New()
	app.SetLogger(log.NewTomLogger())
	app.Rest("/article", &article.Article{})
	app.Start()
}
