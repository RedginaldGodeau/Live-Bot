package route

import (
	"backend/internal/controller"
	"backend/pkg/application"
)

func RouteWS(app *application.App) {
	service := controller.NewControllerServicesFromApp(app)
	wsController := controller.NewWebSocketController(service)

	app.Router.GET("/ws/", wsController.PileLiveShow)
}
