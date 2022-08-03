package main

import (
	"github.com/larbert/mayfly"
	"larbert.top/app"
	"log"
)

func main() {
	e := mayfly.New()

	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "static")

	app.RegisterRoutes(e)

	err := e.Run(":80")
	if err != nil {
		log.Fatalln(err)
	}
}
