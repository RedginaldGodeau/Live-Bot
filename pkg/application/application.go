package application

import (
	"backend/internal/core/stack"
	"backend/pkg/datastore"
	"backend/pkg/discord"
	"backend/pkg/environment"
	"backend/pkg/router"
)

type App struct {
	Env        *environment.Env
	DiscordBot *discord.DiscordBot
	Router     *router.Router
	WS         *router.Router
	DB         *datastore.EntDB
}

func NewApp() (*App, error) {
	env := environment.NewEnv()
	r := router.NewRouter(env.SERVER_WEB_PORT)
	ws := router.NewRouter(env.SERVER_WS_PORT)
	bot, err := discord.NewBot(env.DISCORD_TOKEN)
	if err != nil {
		return nil, err
	}

	database, err := datastore.NewDatabase(env.DATABASE_HOST, env.DATABASE_PORT, env.DATABASE_USER, env.DATABASE_NAME, env.DATABASE_PASSWORD)
	if err != nil {
		return nil, err
	}

	return &App{
		Env:        env,
		DiscordBot: bot,
		Router:     r,
		WS:         ws,
		DB:         database,
	}, nil
}

func (app *App) Start() {
	go app.DiscordBot.Start()
	go stack.StartStack(app.DB)
	app.Router.Serve()
}
