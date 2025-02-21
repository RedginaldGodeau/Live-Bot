ifneq (,$(wildcard .env.local))
	include $(wildcard .env.local)
	export
endif

ENV_FILE := $(wildcard .env)
LOCAL_ENV_FILE := $(ENV_FILE:%=%.local)
DOCKER_COMPOSE := docker compose

_defaut: run

run: $(LOCAL_ENV_FILE)
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml --env-file .env.local up || true

build: $(LOCAL_ENV_FILE)
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up --build

down:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml down -v


ent:
	go run -mod=mod cmd/entc/main.go