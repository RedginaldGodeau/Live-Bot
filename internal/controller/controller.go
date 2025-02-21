package controller

import (
	"backend/pkg/application"
	"backend/pkg/datastore"
	"backend/pkg/environment"
)

type ControllerServices struct {
	environment *environment.Env
	db          *datastore.EntDB
}

func NewControllerServicesFromApp(app *application.App) ControllerServices {
	return ControllerServices{
		environment: app.Env,
		db:          app.DB,
	}
}
