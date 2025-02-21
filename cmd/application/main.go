package main

import (
	"backend/internal/route"
	"backend/pkg/application"
	"log"
)

func main() {
	app, err := application.NewApp()
	if err != nil {
		log.Fatalln(err)
		return
	}

	route.DiscordCommand(app)
	route.Route(app)
	route.RouteWS(app)

	app.Start()
}
