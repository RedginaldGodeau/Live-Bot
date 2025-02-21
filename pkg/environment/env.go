package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DISCORD_APP_ID string
	DISCORD_TOKEN  string

	FRONT_END_URL   string
	SERVER_WEB_PORT string
	SERVER_WS_PORT  string

	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_USER     string
	DATABASE_NAME     string
	DATABASE_PASSWORD string
}

func NewEnv() *Env {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalln(err)
	}

	return &Env{
		DISCORD_APP_ID: os.Getenv("DISCORD_APP_ID"),
		DISCORD_TOKEN:  os.Getenv("DISCORD_TOKEN"),

		FRONT_END_URL:   os.Getenv("FRONT_END_URL"),
		SERVER_WEB_PORT: os.Getenv("SERVER_WEB_PORT"),
		SERVER_WS_PORT:  os.Getenv("SERVER_WS_PORT"),

		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
	}
}
